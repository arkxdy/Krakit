package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
)

type SectionScoresRepository interface {
	UpsertSectionScore(ctx context.Context, sectionScore *db.UpsertSectionScoreParams) error
	GetSectionScores(ctx context.Context, attemptID pgtype.UUID) ([]db.AttemptSectionScore, error)
}

type sectionScoresRepo struct {
	q *db.Queries
}

// GetSectionScores implements [SectionScoresRepository].
func (s *sectionScoresRepo) GetSectionScores(ctx context.Context, attemptID pgtype.UUID) ([]db.AttemptSectionScore, error) {
	attempts, err := s.q.GetSectionScores(ctx, attemptID)
	if err != nil {
		return nil, err
	}
	return attempts, nil
}

// UpsertSectionScore implements [SectionScoresRepository].
func (s *sectionScoresRepo) UpsertSectionScore(ctx context.Context, sectionScore *db.UpsertSectionScoreParams) error {
	err := s.q.UpsertSectionScore(ctx, *sectionScore)
	if err != nil {
		return err
	}
	return nil
}

func NewSectionScoresRepository(q *db.Queries) SectionScoresRepository {
	return &sectionScoresRepo{q: q}
}
