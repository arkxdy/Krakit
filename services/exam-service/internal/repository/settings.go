package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
)

type ExamSettingsRepository interface {
	CreateExamSettings(ctx context.Context, examSettings *db.UpsertExamSettingsParams) (*db.ExamSetting, error)
	PublishExam(ctx context.Context, id pgtype.UUID) error
	GetExamSettings(ctx context.Context, id pgtype.UUID) (*db.ExamSetting, error)
}

type examSettingsRepo struct {
	q *db.Queries
}

// CreateExamSettings implements [ExamSettingsRepository].
func (e *examSettingsRepo) CreateExamSettings(ctx context.Context, examSettings *db.UpsertExamSettingsParams) (*db.ExamSetting, error) {
	setting, err := e.q.UpsertExamSettings(ctx, *examSettings)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

// GetExamSettings implements [ExamSettingsRepository].
func (e *examSettingsRepo) GetExamSettings(ctx context.Context, id pgtype.UUID) (*db.ExamSetting, error) {
	setting, err := e.q.GetExamSettings(ctx, id)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

// PublishExam implements [ExamSettingsRepository].
func (e *examSettingsRepo) PublishExam(ctx context.Context, id pgtype.UUID) error {
	err := e.q.PublishExam(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewExamSettingsRepository(q *db.Queries) ExamSettingsRepository {
	return &examSettingsRepo{q: q}
}
