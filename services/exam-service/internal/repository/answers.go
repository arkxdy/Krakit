package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
)

type AnswerRepository interface {
	CreateAnswer(ctx context.Context, answer *db.UpsertAnswerParams) error
	GetAnswerByAttemptAndQuestion(ctx context.Context, param *db.GetAnswerByAttemptAndQuestionParams) (*db.Answer, error)
	GetAnswersByAttempt(ctx context.Context, attemptID pgtype.UUID) ([]db.Answer, error)
	EvaluateAnswer(ctx context.Context, eval *db.EvaluateAnswerParams) error
}

type answerRepo struct {
	q *db.Queries
}

// CreateAnswer implements [AnswerRepository].
func (a *answerRepo) CreateAnswer(ctx context.Context, answer *db.UpsertAnswerParams) error {
	err := a.q.UpsertAnswer(ctx, *answer)
	if err != nil {
		return err
	}
	return nil
}

// EvaluateAnswer implements [AnswerRepository].
func (a *answerRepo) EvaluateAnswer(ctx context.Context, eval *db.EvaluateAnswerParams) error {
	err := a.q.EvaluateAnswer(ctx, *eval)
	if err != nil {
		return err
	}
	return nil
}

// GetAnswerByAttemptAndQuestion implements [AnswerRepository].
func (a *answerRepo) GetAnswerByAttemptAndQuestion(ctx context.Context, param *db.GetAnswerByAttemptAndQuestionParams) (*db.Answer, error) {
	answer, err := a.q.GetAnswerByAttemptAndQuestion(ctx, *param)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

// GetAnswersByAttempt implements [AnswerRepository].
func (a *answerRepo) GetAnswersByAttempt(ctx context.Context, attemptID pgtype.UUID) ([]db.Answer, error) {
	answers, err := a.q.GetAnswersByAttempt(ctx, attemptID)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func NewAnswerRepository(q *db.Queries) AnswerRepository {
	return &answerRepo{q: q}
}
