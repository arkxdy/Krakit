package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
	"github.com/krakit/exam-service/internal/dto"
	"github.com/krakit/exam-service/internal/repository"
)

type QuestionService interface {
	// QUESTION SET MAPPING
	AttachQuestionSet(ctx context.Context, req dto.AttachQuestionSetReq) error
	GetQuestionSets(ctx context.Context, examID string) ([]dto.AttachQuestionSetReq, error)

	// =========================
	// MONGO (QUESTIONS + PASSAGES)
	// =========================

	CreateQuestions(ctx context.Context, questions []dto.Question) error
	BulkUpsertQuestions(ctx context.Context, questions []dto.Question) error

	CreatePassages(ctx context.Context, passages []dto.Passage) error
	BulkUpsertPassages(ctx context.Context, passages []dto.Passage) error

	DeleteQuestionsBySet(ctx context.Context, setID string) error
	DeletePassagesBySection(ctx context.Context, sectionID string) error
}

type questionService struct {
	question    repository.QuestionRepository
	questionSet repository.QuestionSetsRepository
}

// AttachQuestionSet implements [QuestionService].
func (q *questionService) AttachQuestionSet(ctx context.Context, req dto.AttachQuestionSetReq) error {
	if req.ExamID == "" {
		return fmt.Errorf("questionService.AttachQuestionSet: exam_id is required")
	}
	if req.SectionID == "" {
		return fmt.Errorf("questionService.AttachQuestionSet: section_id is required")
	}
	if req.QuestionSetID == "" {
		return fmt.Errorf("questionService.AttachQuestionSet: question_set_id is required")
	}
	if req.QuestionCount <= 0 {
		return fmt.Errorf("questionService.AttachQuestionSet: question_count must be positive")
	}

	examUUID, err := uuid.Parse(req.ExamID)
	if err != nil {
		return fmt.Errorf("questionService.AttachQuestionSet: invalid exam_id: %w", err)
	}

	sectionUUID, err := uuid.Parse(req.SectionID)
	if err != nil {
		return fmt.Errorf("questionService.AttachQuestionSet: invalid section_id: %w", err)
	}

	// marshal difficulty distribution for JSONB
	diffJSON, err := json.Marshal(req.Difficulty)
	if err != nil {
		return fmt.Errorf("questionService.AttachQuestionSet: marshal difficulty: %w", err)
	}

	_, err = q.questionSet.CreateQuestionSet(ctx, &db.CreateQuestionSetParams{
		ExamID:                 pgtype.UUID{Bytes: examUUID, Valid: true},
		SectionID:              pgtype.UUID{Bytes: sectionUUID, Valid: true},
		QuestionSetID:          req.QuestionSetID,
		QuestionCount:          int32(req.QuestionCount),
		DifficultyDistribution: diffJSON,
	})
	if err != nil {
		return fmt.Errorf("questionService.AttachQuestionSet: %w", err)
	}

	return nil
}

// BulkUpsertPassages implements [QuestionService].
func (q *questionService) BulkUpsertPassages(ctx context.Context, passages []dto.Passage) error {
	if len(passages) == 0 {
		return fmt.Errorf("questionService.BulkUpsertPassages: passages list is empty")
	}

	if err := validatePassages(passages); err != nil {
		return fmt.Errorf("questionService.BulkUpsertPassages: %w", err)
	}

	if err := q.question.BulkUpsertPassages(ctx, passages); err != nil {
		return fmt.Errorf("questionService.BulkUpsertPassages: %w", err)
	}

	return nil
}

// BulkUpsertQuestions implements [QuestionService].
func (q *questionService) BulkUpsertQuestions(ctx context.Context, questions []dto.Question) error {
	if len(questions) == 0 {
		return fmt.Errorf("questionService.BulkUpsertQuestions: questions list is empty")
	}

	if err := validateQuestions(questions); err != nil {
		return fmt.Errorf("questionService.BulkUpsertQuestions: %w", err)
	}

	if err := q.question.BulkUpsertQuestions(ctx, questions); err != nil {
		return fmt.Errorf("questionService.BulkUpsertQuestions: %w", err)
	}

	return nil
}

// CreatePassages implements [QuestionService].
func (q *questionService) CreatePassages(ctx context.Context, passages []dto.Passage) error {
	if len(passages) == 0 {
		return fmt.Errorf("questionService.CreatePassages: passages list is empty")
	}

	if err := validatePassages(passages); err != nil {
		return fmt.Errorf("questionService.CreatePassages: %w", err)
	}

	if err := q.question.InsertPassages(ctx, passages); err != nil {
		return fmt.Errorf("questionService.CreatePassages: %w", err)
	}

	return nil
}

// CreateQuestions implements [QuestionService].
func (q *questionService) CreateQuestions(ctx context.Context, questions []dto.Question) error {
	if len(questions) == 0 {
		return fmt.Errorf("questionService.CreateQuestions: questions list is empty")
	}

	if err := validateQuestions(questions); err != nil {
		return fmt.Errorf("questionService.CreateQuestions: %w", err)
	}

	if err := q.question.InsertQuestions(ctx, questions); err != nil {
		return fmt.Errorf("questionService.CreateQuestions: %w", err)
	}

	return nil
}

// DeletePassagesBySection implements [QuestionService].
func (q *questionService) DeletePassagesBySection(ctx context.Context, sectionID string) error {
	if sectionID == "" {
		return fmt.Errorf("questionService.DeletePassagesBySection: section_id is required")
	}

	if err := q.question.DeletePassagesBySection(ctx, sectionID); err != nil {
		return fmt.Errorf("questionService.DeletePassagesBySection: %w", err)
	}

	return nil
}

// DeleteQuestionsBySet implements [QuestionService].
func (q *questionService) DeleteQuestionsBySet(ctx context.Context, setID string) error {
	if setID == "" {
		return fmt.Errorf("questionService.DeleteQuestionsBySet: set_id is required")
	}

	if err := q.question.DeleteQuestionsBySet(ctx, setID); err != nil {
		return fmt.Errorf("questionService.DeleteQuestionsBySet: %w", err)
	}

	return nil
}

// GetQuestionSets implements [QuestionService].
func (q *questionService) GetQuestionSets(ctx context.Context, examID string) ([]dto.AttachQuestionSetReq, error) {
	if examID == "" {
		return nil, fmt.Errorf("questionService.GetQuestionSets: exam_id is required")
	}

	examUUID, err := uuid.Parse(examID)
	if err != nil {
		return nil, fmt.Errorf("questionService.GetQuestionSets: invalid exam_id: %w", err)
	}

	sets, err := q.questionSet.GetQuestionSetsByExam(ctx, pgtype.UUID{Bytes: examUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("questionService.GetQuestionSets: %w", err)
	}

	result := make([]dto.AttachQuestionSetReq, 0, len(sets))
	for _, s := range sets {
		var difficulty map[string]int
		if err := json.Unmarshal(s.DifficultyDistribution, &difficulty); err != nil {
			return nil, fmt.Errorf("questionService.GetQuestionSets: unmarshal difficulty: %w", err)
		}
		result = append(result, dto.AttachQuestionSetReq{
			ExamID:        s.ExamID.String(),
			SectionID:     s.SectionID.String(),
			QuestionSetID: s.QuestionSetID,
			QuestionCount: int(s.QuestionCount),
			Difficulty:    difficulty,
		})
	}

	return result, nil
}

func NewQuestionService(
	question repository.QuestionRepository,
	questionSet repository.QuestionSetsRepository,
) QuestionService {
	return &questionService{
		question:    question,
		questionSet: questionSet,
	}
}

func validateQuestions(questions []dto.Question) error {
	for i, q := range questions {
		if q.QuestionText == "" {
			return fmt.Errorf("question[%d]: question_text is required", i)
		}
		if q.QuestionSetID == "" {
			return fmt.Errorf("question[%d]: question_set_id is required", i)
		}
		if len(q.Options) == 0 && q.Type != "NUMERIC" {
			return fmt.Errorf("question[%d]: options are required for type %s", i, q.Type)
		}
		if q.Marks <= 0 {
			return fmt.Errorf("question[%d]: marks must be positive", i)
		}
		if q.Type == "" {
			return fmt.Errorf("question[%d]: type is required (MCQ, MSQ, NUMERIC)", i)
		}
	}
	return nil
}

func validatePassages(passages []dto.Passage) error {
	for i, p := range passages {
		if p.PassageText == "" {
			return fmt.Errorf("passage[%d]: passage_text is required", i)
		}
		if p.ExamID == "" {
			return fmt.Errorf("passage[%d]: exam_id is required", i)
		}
		if p.SectionID == "" {
			return fmt.Errorf("passage[%d]: section_id is required", i)
		}
	}
	return nil
}
