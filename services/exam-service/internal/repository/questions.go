package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
)

type QuestionSetsRepository interface {
	CreateQuestionSet(ctx context.Context, questionSet *db.CreateQuestionSetParams) (*db.ExamQuestionSet, error)
	GetQuestionSetsByExam(ctx context.Context, examID pgtype.UUID) ([]db.ExamQuestionSet, error)
	UpdateQuestionSet(ctx context.Context, questionSet *db.UpdateQuestionSetParams) error
	CountQuestionSetsBySection(ctx context.Context, sectionID pgtype.UUID) (int64, error)
}

type questionSetsRepo struct {
	q *db.Queries
}

// CountQuestionSetsBySection implements [QuestionSetsRepository].
func (q *questionSetsRepo) CountQuestionSetsBySection(ctx context.Context, sectionID pgtype.UUID) (int64, error) {
	questionCount, err := q.q.CountQuestionSetsBySection(ctx, sectionID)
	if err != nil {
		return -1, err
	}
	return questionCount, nil
}

// CreateQuestionSet implements [QuestionSetsRepository].
func (q *questionSetsRepo) CreateQuestionSet(ctx context.Context, questionSet *db.CreateQuestionSetParams) (*db.ExamQuestionSet, error) {
	question, err := q.q.CreateQuestionSet(ctx, *questionSet)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

// GetQuestionSetsByExam implements [QuestionSetsRepository].
func (q *questionSetsRepo) GetQuestionSetsByExam(ctx context.Context, examID pgtype.UUID) ([]db.ExamQuestionSet, error) {
	questions, err := q.q.GetQuestionSetsByExam(ctx, examID)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// UpdateQuestionSet implements [QuestionSetsRepository].
func (q *questionSetsRepo) UpdateQuestionSet(ctx context.Context, questionSet *db.UpdateQuestionSetParams) error {
	err := q.q.UpdateQuestionSet(ctx, *questionSet)
	if err != nil {
		return err
	}
	return nil
}

func NewQuestionSetsRepository(q *db.Queries) QuestionSetsRepository {
	return &questionSetsRepo{q: q}
}
