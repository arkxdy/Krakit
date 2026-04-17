package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/auth-service/internal/db/sqlc"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session db.CreateSessionParams) (db.UserSession, error)
	GetSessionByRefreshToken(ctx context.Context, rfrshToken string) (db.UserSession, error)
	ListUserSessions(ctx context.Context, id pgtype.UUID) ([]db.UserSession, error)
	RevokeSession(ctx context.Context, rfrshToken string) error
	RevokeSessionById(ctx context.Context, id pgtype.UUID) error
	RevokeAllUserSessions(ctx context.Context, id pgtype.UUID) error
	DeleteUserSession(ctx context.Context, id pgtype.UUID) error
	DeleteExpiredSessions(ctx context.Context) error
}

type sessionRepo struct {
	q *db.Queries
}

// DeleteUserSession implements [SessionRepository].
func (s *sessionRepo) DeleteUserSession(ctx context.Context, id pgtype.UUID) error {
	err := s.q.DeleteUserSessions(ctx, id)
	if err != nil {
		return fmt.Errorf("delete user session: %w", err)
	}
	return nil
}

// ListUserSessions implements [SessionRepository].
func (s *sessionRepo) ListUserSessions(ctx context.Context, id pgtype.UUID) ([]db.UserSession, error) {
	sessions, err := s.q.ListUserSessions(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("list user session: %w", err)
	}
	return sessions, nil
}

// RevokeAllUserSessions implements [SessionRepository].
func (s *sessionRepo) RevokeAllUserSessions(ctx context.Context, id pgtype.UUID) error {
	err := s.q.RevokeAllUserSessions(ctx, id)
	if err != nil {
		return fmt.Errorf("revoke all sessions: %w", err)
	}
	return nil
}

// RevokeSessionById implements [SessionRepository].
func (s *sessionRepo) RevokeSessionById(ctx context.Context, id pgtype.UUID) error {
	err := s.q.RevokeSessionByID(ctx, id)
	if err != nil {
		return fmt.Errorf("revoke session by id: %w", err)
	}
	return nil
}

// CreateSession implements [SessionRepository].
func (s *sessionRepo) CreateSession(ctx context.Context, session db.CreateSessionParams) (db.UserSession, error) {
	newSession, err := s.q.CreateSession(ctx, session)
	if err != nil {
		return db.UserSession{}, fmt.Errorf("create session: %w", err)
	}
	return newSession, nil
}

// DeleteExpiredSessions implements [SessionRepository].
func (s *sessionRepo) DeleteExpiredSessions(ctx context.Context) error {
	err := s.q.DeleteExpiredSessions(ctx)
	if err != nil {
		return fmt.Errorf("delete expired session: %w", err)
	}
	return nil
}

// GetSessionByRefreshToken implements [SessionRepository].
func (s *sessionRepo) GetSessionByRefreshToken(ctx context.Context, rfrshToken string) (db.UserSession, error) {
	session, err := s.q.GetSessionByRefreshToken(ctx, rfrshToken)
	if err != nil {
		return db.UserSession{}, fmt.Errorf("get session by refresh token: %w", err)
	}
	return session, nil
}

// RevokeSession implements [SessionRepository].
func (s *sessionRepo) RevokeSession(ctx context.Context, rfrshToken string) error {
	err := s.q.RevokeSession(ctx, rfrshToken)
	if err != nil {
		return fmt.Errorf("revoke session: %w", err)
	}
	return nil
}

func NewSessionRepository(q *db.Queries) SessionRepository {
	return &sessionRepo{q: q}
}
