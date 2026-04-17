package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
)

type ExamSectionRepository interface {
	CreateSection(ctx context.Context, section *db.CreateSectionParams) (*db.ExamSection, error)
	GetSectionsByExam(ctx context.Context, examID pgtype.UUID) ([]db.ExamSection, error)
	UpdateSection(ctx context.Context, section *db.UpdateSectionParams) error
}

type sectionRepo struct {
	q *db.Queries
}

// CreateSection implements [ExamSectionRepository].
func (s *sectionRepo) CreateSection(ctx context.Context, section *db.CreateSectionParams) (*db.ExamSection, error) {
	sec, err := s.q.CreateSection(ctx, *section)
	if err != nil {
		return nil, err
	}
	return &sec, nil
}

// GetSectionsByExam implements [ExamSectionRepository].
func (s *sectionRepo) GetSectionsByExam(ctx context.Context, examID pgtype.UUID) ([]db.ExamSection, error) {
	sections, err := s.q.GetSectionsByExam(ctx, examID)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

// UpdateSection implements [ExamSectionRepository].
func (s *sectionRepo) UpdateSection(ctx context.Context, section *db.UpdateSectionParams) error {
	err := s.q.UpdateSection(ctx, *section)
	if err != nil {
		return err
	}
	return nil
}

func NewSectionRepository(q *db.Queries) ExamSectionRepository {
	return &sectionRepo{q: q}
}
