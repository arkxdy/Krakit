package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
)

type ExamRepository interface {
	CreateExam(ctx context.Context, exam *db.CreateExamParams) (*db.Exam, error)
	GetExam(ctx context.Context, id pgtype.UUID) (*db.Exam, error)
	ListExams(ctx context.Context) ([]db.Exam, error)
	ListExamsPaginated(ctx context.Context, args *db.ListExamsPaginatedParams) ([]db.Exam, error)
	ListPublishedExams(ctx context.Context) ([]db.Exam, error)
	DisableExam(ctx context.Context, id pgtype.UUID) error
	ListExamsWithStatus(ctx context.Context) ([]db.ListExamsWithStatusRow, error)
	UpdateExam(ctx context.Context, exam *db.UpdateExamParams) error
	GetExamWithSections(ctx context.Context, id pgtype.UUID) ([]db.GetExamWithSectionsRow, error)
}

type examRepo struct {
	q *db.Queries
}

// CreateExam implements [ExamRepository].
func (e *examRepo) CreateExam(ctx context.Context, exam *db.CreateExamParams) (*db.Exam, error) {
	createdExam, err := e.q.CreateExam(ctx, *exam)
	if err != nil {
		return nil, err
	}
	return &createdExam, nil
}

// DisableExam implements [ExamRepository].
func (e *examRepo) DisableExam(ctx context.Context, id pgtype.UUID) error {
	if err := e.q.DisableExam(ctx, id); err != nil {
		return err
	}
	return nil
}

// GetExam implements [ExamRepository].
func (e *examRepo) GetExam(ctx context.Context, id pgtype.UUID) (*db.Exam, error) {
	exam, err := e.q.GetExam(ctx, id)
	if err != nil {
		return nil, err
	}
	return &exam, nil
}

// GetExamWithSections implements [ExamRepository].
func (e *examRepo) GetExamWithSections(ctx context.Context, id pgtype.UUID) ([]db.GetExamWithSectionsRow, error) {
	exams, err := e.q.GetExamWithSections(ctx, id)
	if err != nil {
		return nil, err
	}
	return exams, nil
}

// ListExams implements [ExamRepository].
func (e *examRepo) ListExams(ctx context.Context) ([]db.Exam, error) {
	exams, err := e.q.ListExams(ctx)
	if err != nil {
		return nil, err
	}
	return exams, nil
}

// ListExamsPaginated implements [ExamRepository].
func (e *examRepo) ListExamsPaginated(ctx context.Context, args *db.ListExamsPaginatedParams) ([]db.Exam, error) {
	exams, err := e.q.ListExamsPaginated(ctx, *args)
	if err != nil {
		return nil, err
	}
	return exams, nil
}

// ListExamsWithStatus implements [ExamRepository].
func (e *examRepo) ListExamsWithStatus(ctx context.Context) ([]db.ListExamsWithStatusRow, error) {
	exams, err := e.q.ListExamsWithStatus(ctx)
	if err != nil {
		return nil, err
	}
	return exams, nil
}

// ListPublishedExams implements [ExamRepository].
func (e *examRepo) ListPublishedExams(ctx context.Context) ([]db.Exam, error) {
	exams, err := e.q.ListPublishedExams(ctx)
	if err != nil {
		return nil, err
	}
	return exams, nil
}

// UpdateExam implements [ExamRepository].
func (e *examRepo) UpdateExam(ctx context.Context, exam *db.UpdateExamParams) error {
	if err := e.q.UpdateExam(ctx, *exam); err != nil {
		return err
	}
	return nil
}

func NewExamRepository(q *db.Queries) ExamRepository {
	return &examRepo{q: q}
}
