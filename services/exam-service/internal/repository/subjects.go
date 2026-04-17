package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
)

type SubjectRepository interface {
	CreateSubject(ctx context.Context, subject *db.CreateSubjectParams) (*db.Subject, error)
	GetSubjects(ctx context.Context) ([]db.Subject, error)
	GetSubject(ctx context.Context, id pgtype.UUID) (*db.Subject, error)
	// UpdateSubject(ctx context.Context, id string, subject *db.Subject) error
	DeleteSubject(ctx context.Context, id pgtype.UUID) error
}

type subjectRepo struct {
	q *db.Queries
}

// CreateSubject implements [SubjectRepository].
func (s *subjectRepo) CreateSubject(ctx context.Context, subject *db.CreateSubjectParams) (*db.Subject, error) {
	sub, err := s.q.CreateSubject(ctx, *subject)
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

// DeleteSubject implements [SubjectRepository].
func (s *subjectRepo) DeleteSubject(ctx context.Context, id pgtype.UUID) error {
	err := s.q.DeleteSubject(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetSubject implements [SubjectRepository].
func (s *subjectRepo) GetSubject(ctx context.Context, id pgtype.UUID) (*db.Subject, error) {
	subject, err := s.q.GetSubject(ctx, id)
	if err != nil {
		return nil, err
	}
	return &subject, nil
}

// GetSubjects implements [SubjectRepository].
func (s *subjectRepo) GetSubjects(ctx context.Context) ([]db.Subject, error) {
	subjects, err := s.q.GetSubjects(ctx)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

func NewSubjectRepository(q *db.Queries) SubjectRepository {
	return &subjectRepo{q: q}
}
