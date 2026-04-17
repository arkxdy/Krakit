package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/repository"
)

type SessionService interface {
	ListSessions(ctx context.Context, userID string, currentJTI string) ([]dto.SessionResponse, error)
	RevokeSession(ctx context.Context, userID string, sessionID string, isAdmin bool) error
}

type sessionService struct {
	session repository.SessionRepository
}

// ListSessions implements [SessionService].
func (s *sessionService) ListSessions(ctx context.Context, userID string, currentJTI string) ([]dto.SessionResponse, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	var currentUUID *uuid.UUID
	if currentJTI != "" {
		if u, err := uuid.Parse(currentJTI); err == nil {
			currentUUID = &u
		}
	}

	sessions, err := s.session.ListUserSessions(ctx, pgtype.UUID{Bytes: userUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("list sessions: %w", err)
	}

	res := make([]dto.SessionResponse, 0, len(sessions))
	for _, sess := range sessions {
		isCurrent := false
		if currentUUID != nil && sess.AccessTokenJti.Valid {
			// AccessTokenJti is stored as the JTI from the access token we generated.
			isCurrent = sess.AccessTokenJti.Bytes == *currentUUID
		}

		createdAt := ""
		if sess.CreatedAt.Time.After(time.Time{}) {
			createdAt = sess.CreatedAt.Time.Format(time.RFC3339)
		}
		expiresAt := ""
		if sess.ExpiresAt.Time.After(time.Time{}) {
			expiresAt = sess.ExpiresAt.Time.Format(time.RFC3339)
		}

		res = append(res, dto.SessionResponse{
			ID:        sess.ID.String(),
			Platform:  string(sess.Platform),
			CreatedAt: createdAt,
			ExpiresAt: expiresAt,
			IsCurrent: isCurrent,
		})
	}

	return res, nil
}

// RevokeSession implements [SessionService].
func (s *sessionService) RevokeSession(ctx context.Context, userID string, sessionID string, isAdmin bool) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user id")
	}

	sessionUUID, err := uuid.Parse(sessionID)
	if err != nil {
		return errors.New("invalid session id")
	}

	// Non-admins can only revoke their own active sessions.
	if !isAdmin {
		sessions, err := s.session.ListUserSessions(ctx, pgtype.UUID{Bytes: userUUID, Valid: true})
		if err != nil {
			return fmt.Errorf("list sessions: %w", err)
		}

		owned := false
		for _, sess := range sessions {
			if sess.ID.Valid && sess.ID.Bytes == sessionUUID {
				owned = true
				break
			}
		}
		if !owned {
			return errors.New("session not found")
		}
	}

	return s.session.RevokeSessionById(ctx, pgtype.UUID{Bytes: sessionUUID, Valid: true})
}

func NewSessionService(r repository.SessionRepository) SessionService {
	return &sessionService{session: r}
}
