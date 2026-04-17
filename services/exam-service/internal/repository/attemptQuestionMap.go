package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
)

type AttemptQuestionMapRepository interface {
	BulkInsertAttemptMap(ctx context.Context, params []db.BulkInsertAttemptMapParams) (int64, error)
	GetAttemptQuestions(ctx context.Context, attemptID pgtype.UUID) ([]db.AttemptQuestionMap, error)
	GetAttemptQuestionsBySection(ctx context.Context, params *db.GetAttemptQuestionsBySectionParams) ([]db.AttemptQuestionMap, error)
}

type attemptQuestionMapRepo struct {
	q *db.Queries
}

// BulkInsertAttemptMap implements [AttemptQuestionMapRepository].
func (a *attemptQuestionMapRepo) BulkInsertAttemptMap(ctx context.Context, params []db.BulkInsertAttemptMapParams) (int64, error) {
	success, err := a.q.BulkInsertAttemptMap(ctx, params)
	if err != nil {
		return -1, err
	}
	return success, nil
}

// GetAttemptQuestions implements [AttemptQuestionMapRepository].
func (a *attemptQuestionMapRepo) GetAttemptQuestions(ctx context.Context, attemptID pgtype.UUID) ([]db.AttemptQuestionMap, error) {
	questions, err := a.q.GetAttemptQuestions(ctx, attemptID)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// GetAttemptQuestionsBySection implements [AttemptQuestionMapRepository].
func (a *attemptQuestionMapRepo) GetAttemptQuestionsBySection(ctx context.Context, params *db.GetAttemptQuestionsBySectionParams) ([]db.AttemptQuestionMap, error) {
	questions, err := a.q.GetAttemptQuestionsBySection(ctx, *params)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func NewAttemptQuestionMapRepository(q *db.Queries) AttemptQuestionMapRepository {
	return &attemptQuestionMapRepo{q: q}
}
