package service

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/auth-service/internal/db/sqlc"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/repository"
	"github.com/krakit/auth-service/internal/utils"
)

type UserService interface {
	GetCurrentUser(ctx context.Context, userID string) (dto.UserResponse, error)
	UpdateProfile(ctx context.Context, userID string, req dto.UpdateProfileRequest) (dto.UserResponse, error)
	ChangePassword(ctx context.Context, userID string, req dto.ChangePasswordRequest) error
}

type userService struct {
	user repository.UserRepository
}

// ChangePassword implements [UserService].
func (u *userService) ChangePassword(ctx context.Context, userID string, req dto.ChangePasswordRequest) error {
	if req.NewPassword == req.OldPassword {
		return errors.New("new password cannot be same")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user id")
	}

	user, err := u.user.GetUserByID(ctx, pgtype.UUID{Bytes: userUUID, Valid: true})
	if err != nil {
		return errors.New("user not found")
	}

	if err := utils.CheckPassword(req.OldPassword, user.PasswordHash); err != nil {
		return errors.New("invalid old password")
	}

	newHashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("failed to change password")
	}

	return u.user.UpdateUserPassword(ctx, db.UpdateUserPasswordParams{ID: pgtype.UUID{Bytes: userUUID, Valid: true}, PasswordHash: newHashedPassword})
}

// GetCurrentUser implements [UserService].
func (u *userService) GetCurrentUser(ctx context.Context, userID string) (dto.UserResponse, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return dto.UserResponse{}, errors.New("invalid user id")
	}

	user, err := u.user.GetUserByID(ctx, pgtype.UUID{Bytes: userUUID, Valid: true})
	if err != nil {
		return dto.UserResponse{}, errors.New("failed to get user")
	}

	return dto.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		FullName:  user.FullName.String,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Role:      string(user.Role),
		Plan:      string(user.Plan),
	}, nil
}

// UpdateProfile implements [UserService].
func (u *userService) UpdateProfile(ctx context.Context, userID string, req dto.UpdateProfileRequest) (dto.UserResponse, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return dto.UserResponse{}, errors.New("invalid user id")
	}

	req.FirstName = strings.TrimSpace(req.FirstName)
	req.LastName = strings.TrimSpace(req.LastName)
	fullName := strings.TrimSpace(req.FirstName + " " + req.LastName)

	user, err := u.user.UpdateUserProfile(ctx, db.UpdateUserProfileParams{
		ID:        pgtype.UUID{Bytes: userUUID, Valid: true},
		FirstName: pgtype.Text{String: req.FirstName, Valid: true},
		LastName:  pgtype.Text{String: req.LastName, Valid: true},
		FullName:  pgtype.Text{String: fullName, Valid: true},
	})
	if err != nil {
		return dto.UserResponse{}, errors.New("failed to update profile")
	}

	return dto.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		FullName:  user.FullName.String,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Role:      string(user.Role),
		Plan:      string(user.Plan),
	}, nil
}

func NewUserService(u repository.UserRepository) UserService {
	return &userService{user: u}
}
