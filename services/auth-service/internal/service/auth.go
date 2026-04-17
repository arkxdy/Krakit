package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krakit/auth-service/internal/cache"
	db "github.com/krakit/auth-service/internal/db/sqlc"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/repository"
	"github.com/krakit/auth-service/internal/utils"
	"google.golang.org/api/idtoken"
)

type AuthService interface {
	Signup(ctx context.Context, req dto.SignupRequest) (dto.AuthResponse, string, error)
	Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, string, error)
	Refresh(ctx context.Context, refreshToken string) (dto.AuthResponse, string, error)
	Logout(ctx context.Context, refreshToken string) error
	LogoutAll(ctx context.Context, userID string) error
	GoogleLogin(ctx context.Context, idToken string) (dto.AuthResponse, string, error)
	createSessionAndTokenTx(ctx context.Context, q *db.Queries, user db.User, platform db.PlatformType) (dto.AuthResponse, string, error)
	withTx(ctx context.Context, fn func(q *db.Queries) error) error
}

type authService struct {
	db         *pgxpool.Pool
	q          *db.Queries
	user       repository.UserRepository
	session    repository.SessionRepository
	permission repository.PermissionRepository
	jwtMaker   *utils.JWTMaker
	config     *utils.Config
	cache      cache.Cache
}

// GoogleLogin implements [AuthService].
func (a *authService) GoogleLogin(ctx context.Context, idToken string) (dto.AuthResponse, string, error) {
	payload, err := idtoken.Validate(ctx, idToken, a.config.GoogleClientID)
	if err != nil {
		return dto.AuthResponse{}, "", errors.New("invalid google token")
	}

	// SAFE extraction
	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)

	if email == "" {
		return dto.AuthResponse{}, "", errors.New("email not found in google token")
	}

	var resp dto.AuthResponse
	var refreshToken string

	err = a.withTx(ctx, func(q *db.Queries) error {

		userRepo := repository.NewUserRespository(q)

		// Try get user
		user, err := userRepo.GetUserByEmail(ctx, email)

		if err != nil {
			// create if not found
			user, err = userRepo.CreateUser(ctx, db.CreateUserParams{
				Email:        email,
				PasswordHash: "", // OAuth user
				FullName:     pgtype.Text{String: name, Valid: true},
				Role:         db.RoleTypeCandidate,
				Plan:         db.UserPlanTypeFree,
			})
			if err != nil {
				if isUniqueViolation(err) {
					// race condition fallback
					user, err = userRepo.GetUserByEmail(ctx, email)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}
		}

		// reuse core logic
		resp, refreshToken, err = a.createSessionAndTokenTx(
			ctx,
			q,
			user,
			db.PlatformTypeWeb,
		)

		return err
	})

	if err != nil {
		return dto.AuthResponse{}, "", err
	}

	return resp, refreshToken, nil
}

// Login implements [AuthService].
func (a *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, string, error) {
	var resp dto.AuthResponse
	var refreshToken string

	email := strings.ToLower(strings.TrimSpace(req.Email))

	err := a.withTx(ctx, func(q *db.Queries) error {

		userRepo := repository.NewUserRespository(q)

		user, err := userRepo.GetUserByEmail(ctx, email)
		if err != nil {
			return errors.New("invalid credentials")
		}

		// check password
		if err := utils.CheckPassword(req.Password, user.PasswordHash); err != nil {
			return errors.New("invalid credentials")
		}

		// reuse core logic
		resp, refreshToken, err = a.createSessionAndTokenTx(
			ctx,
			q,
			user,
			db.PlatformType(req.Platform),
		)

		return err
	})

	if err != nil {
		return dto.AuthResponse{}, "", err
	}

	return resp, refreshToken, nil
}

// Logout implements [AuthService].
func (a *authService) Logout(ctx context.Context, refreshToken string) error {
	hashed := utils.HashToken(refreshToken, a.config.TokenSecret)

	err := a.session.RevokeSession(ctx, hashed)
	if err != nil {
		return errors.New("failed to logout")
	}

	return nil
}

// LogoutAll implements [AuthService].
func (a *authService) LogoutAll(ctx context.Context, userID string) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("failed to logout-all")
	}
	err = a.session.RevokeAllUserSessions(ctx, pgtype.UUID{Bytes: userUUID, Valid: true})
	if err != nil {
		return errors.New("failed to logout all sessions")
	}
	return nil
}

// Refresh implements [AuthService].
func (a *authService) Refresh(ctx context.Context, refreshToken string) (dto.AuthResponse, string, error) {
	var resp dto.AuthResponse
	var newRefreshToken string

	hashed := utils.HashToken(refreshToken, a.config.TokenSecret)

	err := a.withTx(ctx, func(q *db.Queries) error {

		sessionRepo := repository.NewSessionRepository(q)
		userRepo := repository.NewUserRespository(q)

		// 1. Get session
		session, err := sessionRepo.GetSessionByRefreshToken(ctx, hashed)
		if err != nil {
			return errors.New("invalid refresh token")
		}

		// 2. Check expiry
		if session.ExpiresAt.Time.Before(time.Now()) {
			return errors.New("refresh token expired")
		}

		// 3. Revoke old session (rotation)
		err = sessionRepo.RevokeSession(ctx, hashed)
		if err != nil {
			return err
		}

		// 4. Get user
		user, err := userRepo.GetUserByID(ctx, session.UserID)
		if err != nil {
			return err
		}

		// 5. Create new session + tokens
		resp, newRefreshToken, err = a.createSessionAndTokenTx(
			ctx,
			q,
			user,
			session.Platform,
		)

		return err
	})

	if err != nil {
		return dto.AuthResponse{}, "", err
	}

	return resp, newRefreshToken, nil
}

// Signup implements [AuthService].
func (a *authService) Signup(ctx context.Context, req dto.SignupRequest) (dto.AuthResponse, string, error) {

	var resp dto.AuthResponse
	var refreshToken string

	// 1. Normalize input
	email := strings.ToLower(strings.TrimSpace(req.Email))
	fullName := strings.TrimSpace(req.FirstName + " " + req.LastName)

	// 2. Hash password
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return resp, "", fmt.Errorf("hash password: %w", err)
	}

	// 3. Run everything in transaction
	err = a.withTx(ctx, func(q *db.Queries) error {

		// Create repos bound to transaction
		userRepo := repository.NewUserRespository(q)

		// 4. Create user
		user, err := userRepo.CreateUser(ctx, db.CreateUserParams{
			Email:        email,
			PasswordHash: hash,
			FirstName:    pgtype.Text{String: req.FirstName, Valid: true},
			LastName:     pgtype.Text{String: req.LastName, Valid: true},
			FullName:     pgtype.Text{String: fullName, Valid: true},
			Role:         db.RoleTypeCandidate,
			Plan:         db.UserPlanTypeFree,
		})
		if err != nil {
			// handle duplicate email (optional helper)
			if isUniqueViolation(err) {
				return fmt.Errorf("user already exists: %w", err)
			}
			return fmt.Errorf("create user: %w", err)
		}

		// 5. Create session + tokens
		resp, refreshToken, err = a.createSessionAndTokenTx(
			ctx,
			q,
			user,
			db.PlatformType(req.Platform),
		)
		if err != nil {
			return fmt.Errorf("create session/token: %w", err)
		}

		return nil
	})

	// 6. Handle transaction error
	if err != nil {
		return dto.AuthResponse{}, "", err
	}

	// 7. Return response
	return resp, refreshToken, nil
}

func (a *authService) createSessionAndTokenTx(ctx context.Context, q *db.Queries, user db.User, platform db.PlatformType) (dto.AuthResponse, string, error) {

	// repos bound to transaction
	sessionRepo := repository.NewSessionRepository(q)
	permissionRepo := repository.NewPerissionRepository(q)

	// 1. Get permissions
	permissions, err := permissionRepo.GetPermissionsByRole(ctx, user.Role)
	if err != nil {
		return dto.AuthResponse{}, "", fmt.Errorf("get permissions: %w", err)
	}

	perms := make([]string, 0, len(permissions))
	for _, p := range permissions {
		perms = append(perms, p.Name)
	}

	// 2. Generate access token
	accessToken, jti, err := a.jwtMaker.GenerateToken(
		user.ID.String(),
		user.Email,
		string(user.Role),
		string(user.Plan),
		perms,
	)
	if err != nil {
		return dto.AuthResponse{}, "", fmt.Errorf("generate token: %w", err)
	}

	// 3. Generate refresh token
	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return dto.AuthResponse{}, "", fmt.Errorf("generate refresh token: %w", err)
	}

	// 4. Hash refresh token (NEVER store raw)
	hashed := utils.HashToken(refreshToken, a.config.TokenSecret)

	// 5. Parse JTI
	jtiUUID, err := uuid.Parse(jti)
	if err != nil {
		return dto.AuthResponse{}, "", fmt.Errorf("parse jti: %w", err)
	}

	// 6. Create session
	_, err = sessionRepo.CreateSession(ctx, db.CreateSessionParams{
		UserID:         user.ID,
		Platform:       platform,
		RefreshToken:   hashed,
		AccessTokenJti: pgtype.UUID{Bytes: jtiUUID, Valid: true},
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(time.Duration(a.config.TokenDuration) * time.Second),
			Valid: true,
		},
	})
	if err != nil {
		return dto.AuthResponse{}, "", fmt.Errorf("create session: %w", err)
	}

	// 7. Build response
	return dto.AuthResponse{
		User: dto.AuthUser{
			ID:       user.ID.String(),
			Email:    user.Email,
			FullName: user.FullName.String,
			Role:     string(user.Role),
			Plan:     string(user.Plan),
		},
		AccessToken: accessToken,
		ExpiresIn:   int64(a.config.TokenDuration),
	}, refreshToken, nil
}

func (a *authService) withTx(ctx context.Context, fn func(q *db.Queries) error) error {

	tx, err := a.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := a.q.WithTx(tx)

	if err := fn(qtx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func NewAuthService(
	u repository.UserRepository,
	s repository.SessionRepository,
	p repository.PermissionRepository,
	j *utils.JWTMaker,
	g *utils.Config,
	c cache.Cache,
) AuthService {
	return &authService{user: u, session: s, permission: p, jwtMaker: j, config: g, cache: c}
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" // unique_violation
	}
	return false
}
