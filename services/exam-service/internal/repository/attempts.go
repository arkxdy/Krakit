package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
)

type AttemptRepository interface {
	CreateAttempt(ctx context.Context, attempt *db.CreateAttemptParams) (*db.Attempt, error)
	GetAttempt(ctx context.Context, id pgtype.UUID) (*db.Attempt, error)
	GetActiveAttempt(ctx context.Context, args *db.GetActiveAttemptParams) (*db.Attempt, error)
	CompleteAttempt(ctx context.Context, args *db.CompleteAttemptParams) (*db.Attempt, error)
	GetLatestAttempt(ctx context.Context, args *db.GetLatestAttemptParams) (*db.Attempt, error)
	GetAttemptsByUser(ctx context.Context, userID pgtype.UUID) ([]db.Attempt, error)
	GetAttemptsByExam(ctx context.Context, examID pgtype.UUID) ([]db.Attempt, error)
	GetAttemptsByExamPaginated(ctx context.Context, args *db.GetAttemptsByExamPaginatedParams) ([]db.Attempt, error)
	GetAttemptStatus(ctx context.Context, id pgtype.UUID) (pgtype.Text, error)
	GetAttemptResult(ctx context.Context, id pgtype.UUID) (*db.GetAttemptResultRow, error)
	GetAverageScoreByExam(ctx context.Context, examID pgtype.UUID) (float64, error)
	CountAttemptsByExam(ctx context.Context, examID pgtype.UUID) (int64, error)
}

type attemptRepo struct {
	q *db.Queries
}

// CompleteAttempt implements [AttemptRepository].
func (a *attemptRepo) CompleteAttempt(ctx context.Context, args *db.CompleteAttemptParams) (*db.Attempt, error) {
	attempt, err := a.q.CompleteAttempt(ctx, *args)
	if err != nil {
		return nil, err
	}
	return &attempt, nil
}

// CountAttemptsByExam implements [AttemptRepository].
func (a *attemptRepo) CountAttemptsByExam(ctx context.Context, examID pgtype.UUID) (int64, error) {
	attempts, err := a.q.CountAttemptsByExam(ctx, examID)
	if err != nil {
		return -1, err
	}
	return attempts, nil
}

// CreateAttempt implements [AttemptRepository].
func (a *attemptRepo) CreateAttempt(ctx context.Context, attempt *db.CreateAttemptParams) (*db.Attempt, error) {
	newAttempt, err := a.q.CreateAttempt(ctx, *attempt)
	if err != nil {
		return nil, err
	}
	return &newAttempt, nil
}

// GetActiveAttempt implements [AttemptRepository].
func (a *attemptRepo) GetActiveAttempt(ctx context.Context, args *db.GetActiveAttemptParams) (*db.Attempt, error) {
	attempt, err := a.q.GetActiveAttempt(ctx, *args)
	if err != nil {
		return nil, err
	}
	return &attempt, nil
}

// GetAttempt implements [AttemptRepository].
func (a *attemptRepo) GetAttempt(ctx context.Context, id pgtype.UUID) (*db.Attempt, error) {
	attempt, err := a.q.GetAttempt(ctx, id)
	if err != nil {
		return nil, err
	}
	return &attempt, nil
}

// GetAttemptResult implements [AttemptRepository].
func (a *attemptRepo) GetAttemptResult(ctx context.Context, id pgtype.UUID) (*db.GetAttemptResultRow, error) {
	attempt, err := a.q.GetAttemptResult(ctx, id)
	if err != nil {
		return nil, err
	}
	return &attempt, nil
}

// GetAttemptStatus implements [AttemptRepository].
func (a *attemptRepo) GetAttemptStatus(ctx context.Context, id pgtype.UUID) (pgtype.Text, error) {
	status, err := a.q.GetAttemptStatus(ctx, id)
	if err != nil {
		return pgtype.Text{}, err
	}
	switch v := status.(type) {
	case pgtype.Text:
		return v, nil
	case string:
		return pgtype.Text{String: v, Valid: true}, nil
	case []byte:
		return pgtype.Text{String: string(v), Valid: true}, nil
	default:
		return pgtype.Text{}, fmt.Errorf("unexpected attempt status type %T", status)
	}
}

// GetAttemptsByExam implements [AttemptRepository].
func (a *attemptRepo) GetAttemptsByExam(ctx context.Context, examID pgtype.UUID) ([]db.Attempt, error) {
	attempts, err := a.q.GetAttemptsByExam(ctx, examID)
	if err != nil {
		return nil, err
	}
	return attempts, nil
}

// GetAttemptsByExamPaginated implements [AttemptRepository].
func (a *attemptRepo) GetAttemptsByExamPaginated(ctx context.Context, args *db.GetAttemptsByExamPaginatedParams) ([]db.Attempt, error) {
	attempts, err := a.q.GetAttemptsByExamPaginated(ctx, *args)
	if err != nil {
		return nil, err
	}
	return attempts, nil
}

// GetAttemptsByUser implements [AttemptRepository].
func (a *attemptRepo) GetAttemptsByUser(ctx context.Context, userID pgtype.UUID) ([]db.Attempt, error) {
	attempts, err := a.q.GetAttemptsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return attempts, nil
}

// GetAverageScoreByExam implements [AttemptRepository].
func (a *attemptRepo) GetAverageScoreByExam(ctx context.Context, examID pgtype.UUID) (float64, error) {
	score, err := a.q.GetAverageScoreByExam(ctx, examID)
	if err != nil {
		return -1, err
	}

	switch v := score.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case []byte:
		parsed, parseErr := strconv.ParseFloat(string(v), 64)
		if parseErr != nil {
			return -1, fmt.Errorf("parse average score from []byte: %w", parseErr)
		}
		return parsed, nil
	case string:
		parsed, parseErr := strconv.ParseFloat(v, 64)
		if parseErr != nil {
			return -1, fmt.Errorf("parse average score from string: %w", parseErr)
		}
		return parsed, nil
	default:
		return -1, fmt.Errorf("unexpected average score type %T", score)
	}
}

// GetLatestAttempt implements [AttemptRepository].
func (a *attemptRepo) GetLatestAttempt(ctx context.Context, args *db.GetLatestAttemptParams) (*db.Attempt, error) {
	attempt, err := a.q.GetLatestAttempt(ctx, *args)
	if err != nil {
		return nil, err
	}
	return &attempt, nil
}

func NewAttemptRepository(q *db.Queries) AttemptRepository {
	return &attemptRepo{q: q}
}
