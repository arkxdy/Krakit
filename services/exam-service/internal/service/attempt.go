package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
	"github.com/krakit/exam-service/internal/dto"
	"github.com/krakit/exam-service/internal/repository"
	"github.com/krakit/exam-service/internal/worker"
)

type AttemptService interface {
	// =====================================================
	// USER APIs
	// =====================================================

	StartAttempt(ctx context.Context, userID string, req dto.StartAttemptReq) (*dto.Attempt, error)

	GetAttempt(ctx context.Context, attemptID string) (*dto.Attempt, error)

	GetAttemptQuestions(ctx context.Context, attemptID string) ([]dto.QuestionWithMeta, error)

	GetAttemptQuestionsBySection(ctx context.Context, attemptID, sectionID string) ([]dto.QuestionWithMeta, error)

	SaveAnswer(ctx context.Context, ans dto.Answer) error

	GetAnswers(ctx context.Context, attemptID string) ([]dto.Answer, error)

	SubmitAttempt(ctx context.Context, attemptID string) (*dto.AttemptResult, error)

	GetResult(ctx context.Context, attemptID string) (*dto.AttemptResult, error)

	GetUserAttempts(ctx context.Context, userID string) ([]dto.Attempt, error)

	// =====================================================
	// INTERNAL (CRITICAL FLOW)
	// =====================================================

	// Build attempt question map (PG + Mongo)
	PrepareAttempt(ctx context.Context, attemptID string) error

	// Evaluation worker (Kafka consumer)
	EvaluateAttempt(ctx context.Context, attemptID string) error
}

type attemptService struct {
	attempt        repository.AttemptRepository
	attemptQuesMap repository.AttemptQuestionMapRepository
	answer         repository.AnswerRepository
	question       repository.QuestionRepository
	questionSet    repository.QuestionSetsRepository
	score          repository.SectionScoresRepository
	pool           *worker.Pool
}

// EvaluateAttempt implements [AttemptService].
func (a *attemptService) EvaluateAttempt(ctx context.Context, attemptID string) error {
	if attemptID == "" {
		return fmt.Errorf("attemptService.EvaluateAttempt: attempt_id is required")
	}

	attemptUUID, err := uuid.Parse(attemptID)
	if err != nil {
		return fmt.Errorf("attemptService.EvaluateAttempt: invalid attempt_id: %w", err)
	}

	// 1. get all answers submitted
	answers, err := a.answer.GetAnswersByAttempt(ctx, pgtype.UUID{Bytes: attemptUUID, Valid: true})
	if err != nil {
		return fmt.Errorf("attemptService.EvaluateAttempt: get answers: %w", err)
	}

	// 2. fetch question content from Mongo for each answer (has correct_answer, marks, negative_marks)
	questionIDs := make([]string, 0, len(answers))
	for _, ans := range answers {
		questionIDs = append(questionIDs, ans.QuestionID)
	}

	questions, err := a.question.GetQuestionsByIDs(ctx, questionIDs)
	if err != nil {
		return fmt.Errorf("attemptService.EvaluateAttempt: get questions: %w", err)
	}

	// index questions by ID for O(1) lookup
	questionMap := make(map[string]dto.Question, len(questions))
	for _, q := range questions {
		questionMap[q.ID] = q
	}

	// 3. score each answer, accumulate section totals
	type sectionAccum struct {
		score   float64
		correct int
		wrong   int
	}
	sectionTotals := make(map[string]*sectionAccum)
	totalScore := 0.0

	for _, ans := range answers {
		q, ok := questionMap[ans.QuestionID]
		if !ok {
			// question missing from Mongo — skip, log
			log.Printf("attemptService.EvaluateAttempt: question %s not found in Mongo, skipping", ans.QuestionID)
			continue
		}

		var selected interface{}
		if err := json.Unmarshal(ans.SelectedAnswer, &selected); err != nil {
			return fmt.Errorf("attemptService.EvaluateAttempt: unmarshal answer for question %s: %w", ans.QuestionID, err)
		}

		isCorrect := isAnswerCorrect(q, selected)
		marksAwarded := calculateMarks(q, isCorrect)
		totalScore += marksAwarded

		// update answer row with evaluation result
		questionUUID, _ := uuid.Parse(ans.QuestionID)
		if err := a.answer.EvaluateAnswer(ctx, &db.EvaluateAnswerParams{
			AttemptID:    pgtype.UUID{Bytes: attemptUUID, Valid: true},
			QuestionID:   questionUUID.String(),
			IsCorrect:    pgtype.Bool{Bool: isCorrect, Valid: true},
			MarksAwarded: pgtype.Float8{Float64: marksAwarded, Valid: true},
		}); err != nil {
			return fmt.Errorf("attemptService.EvaluateAttempt: evaluate answer: %w", err)
		}

		// accumulate per section
		if _, exists := sectionTotals[q.SectionID]; !exists {
			sectionTotals[q.SectionID] = &sectionAccum{}
		}
		acc := sectionTotals[q.SectionID]
		acc.score += marksAwarded
		if isCorrect {
			acc.correct++
		} else if selected != nil {
			acc.wrong++ // only count wrong if answered — unanswered don't count as wrong
		}
	}

	// 4. upsert section scores
	for sectionID, acc := range sectionTotals {
		sectionUUID, err := uuid.Parse(sectionID)
		if err != nil {
			return fmt.Errorf("attemptService.EvaluateAttempt: invalid section_id %s: %w", sectionID, err)
		}
		if err := a.score.UpsertSectionScore(ctx, &db.UpsertSectionScoreParams{
			AttemptID: pgtype.UUID{Bytes: attemptUUID, Valid: true},
			SectionID: pgtype.UUID{Bytes: sectionUUID, Valid: true},
			Score:     pgtype.Float8{Float64: acc.score, Valid: true},
			Correct:   pgtype.Int4{Int32: int32(acc.correct), Valid: true},
			Wrong:     pgtype.Int4{Int32: int32(acc.wrong), Valid: true},
		}); err != nil {
			return fmt.Errorf("attemptService.EvaluateAttempt: upsert section score: %w", err)
		}
	}

	// 5. complete attempt — marks it submitted with final score
	if _, err := a.attempt.CompleteAttempt(ctx, &db.CompleteAttemptParams{
		ID:         pgtype.UUID{Bytes: attemptUUID, Valid: true},
		TotalScore: pgtype.Float8{Float64: totalScore, Valid: true},
	}); err != nil {
		return fmt.Errorf("attemptService.EvaluateAttempt: complete attempt: %w", err)
	}

	return nil
}

// GetAnswers implements [AttemptService].
func (a *attemptService) GetAnswers(ctx context.Context, attemptID string) ([]dto.Answer, error) {
	if attemptID == "" {
		return nil, fmt.Errorf("examService.GetAnswers: attempt_id is required")
	}

	attemptUUID, err := uuid.Parse(attemptID)
	if err != nil {
		return nil, fmt.Errorf("examService.GetAnswers: invalid attempt_id: %w", err)
	}

	answers, err := a.answer.GetAnswersByAttempt(ctx, pgtype.UUID{Bytes: attemptUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("examService.GetAnswers: %w", err)
	}

	result := make([]dto.Answer, 0, len(answers))
	for _, a := range answers {
		var selected interface{}
		if err := json.Unmarshal(a.SelectedAnswer, &selected); err != nil {
			return nil, fmt.Errorf("examService.GetAnswers: unmarshal selected_answer: %w", err)
		}
		result = append(result, dto.Answer{
			AttemptID:      a.AttemptID.String(),
			QuestionID:     a.QuestionID,
			SelectedAnswer: selected,
		})
	}

	return result, nil
}

// GetAttempt implements [AttemptService].
func (a *attemptService) GetAttempt(ctx context.Context, attemptID string) (*dto.Attempt, error) {
	if attemptID == "" {
		return nil, fmt.Errorf("examService.GetAttempt: attempt_id is required")
	}

	attemptUUID, err := uuid.Parse(attemptID)
	if err != nil {
		return nil, fmt.Errorf("examService.GetAttempt: invalid attempt_id: %w", err)
	}

	attempt, err := a.attempt.GetAttempt(ctx, pgtype.UUID{Bytes: attemptUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("examService.GetAttempt: %w", err)
	}

	status, ok := attempt.Status.(string)
	if !ok {
		return nil, fmt.Errorf("attemptService.GetAttempt: unexpected Status type %T", attempt.Status)
	}

	return &dto.Attempt{
		ID:     attempt.ID.String(),
		UserID: attempt.UserID.String(),
		ExamID: attempt.ExamID.String(),
		Status: status,
		Score:  attempt.TotalScore.Float64,
	}, nil
}

// GetAttemptQuestions implements [AttemptService].
func (a *attemptService) GetAttemptQuestions(ctx context.Context, attemptID string) ([]dto.QuestionWithMeta, error) {
	if attemptID == "" {
		return nil, fmt.Errorf("attemptService.GetAttemptQuestions: attempt_id is required")
	}

	attemptUUID, err := uuid.Parse(attemptID)
	if err != nil {
		return nil, fmt.Errorf("attemptService.GetAttemptQuestions: invalid attempt_id: %w", err)
	}

	// 1. get ordered question map from PG
	mapped, err := a.attemptQuesMap.GetAttemptQuestions(ctx, pgtype.UUID{Bytes: attemptUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("attemptService.GetAttemptQuestions: %w", err)
	}

	return a.hydrateQuestions(ctx, mapped)
}

// GetAttemptQuestionsBySection implements [AttemptService].
func (a *attemptService) GetAttemptQuestionsBySection(ctx context.Context, attemptID string, sectionID string) ([]dto.QuestionWithMeta, error) {
	if attemptID == "" {
		return nil, fmt.Errorf("attemptService.GetAttemptQuestionsBySection: attempt_id is required")
	}
	if sectionID == "" {
		return nil, fmt.Errorf("attemptService.GetAttemptQuestionsBySection: section_id is required")
	}

	attemptUUID, err := uuid.Parse(attemptID)
	if err != nil {
		return nil, fmt.Errorf("attemptService.GetAttemptQuestionsBySection: invalid attempt_id: %w", err)
	}

	sectionUUID, err := uuid.Parse(sectionID)
	if err != nil {
		return nil, fmt.Errorf("attemptService.GetAttemptQuestionsBySection: invalid section_id: %w", err)
	}

	mapped, err := a.attemptQuesMap.GetAttemptQuestionsBySection(ctx, &db.GetAttemptQuestionsBySectionParams{
		AttemptID: pgtype.UUID{Bytes: attemptUUID, Valid: true},
		SectionID: pgtype.UUID{Bytes: sectionUUID, Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("attemptService.GetAttemptQuestionsBySection: %w", err)
	}

	return a.hydrateQuestions(ctx, mapped)
}

// GetResult implements [AttemptService].
func (a *attemptService) GetResult(ctx context.Context, attemptID string) (*dto.AttemptResult, error) {
	if attemptID == "" {
		return nil, fmt.Errorf("examService.GetResult: attempt_id is required")
	}

	attemptUUID, err := uuid.Parse(attemptID)
	if err != nil {
		return nil, fmt.Errorf("examService.GetResult: invalid attempt_id: %w", err)
	}

	attempt, err := a.attempt.GetAttemptResult(ctx, pgtype.UUID{Bytes: attemptUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("examService.GetResult: %w", err)
	}

	sectionScores, err := a.score.GetSectionScores(ctx, pgtype.UUID{Bytes: attemptUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("examService.GetResult: section scores: %w", err)
	}

	sections := make([]dto.SectionResult, 0, len(sectionScores))
	for _, s := range sectionScores {
		sections = append(sections, dto.SectionResult{
			SectionID: s.SectionID.String(),
			Score:     s.Score.Float64,
			Correct:   int(s.Correct.Int32),
			Wrong:     int(s.Wrong.Int32),
		})
	}

	return &dto.AttemptResult{
		AttemptID: attempt.ID.String(),
		Score:     attempt.TotalScore.Float64,
		Sections:  sections,
	}, nil
}

// GetUserAttempts implements [AttemptService].
func (a *attemptService) GetUserAttempts(ctx context.Context, userID string) ([]dto.Attempt, error) {
	if userID == "" {
		return nil, fmt.Errorf("examService.GetUserAttempts: user_id is required")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("examService.GetUserAttempts: invalid user_id: %w", err)
	}

	attempts, err := a.attempt.GetAttemptsByUser(ctx, pgtype.UUID{Bytes: userUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("examService.GetUserAttempts: %w", err)
	}

	result := make([]dto.Attempt, 0, len(attempts))
	for _, a := range attempts {
		status, ok := a.Status.(string)
		if !ok {
			return nil, fmt.Errorf("attemptService.StartAttempt: unexpected Status type %T", a.Status)
		}
		result = append(result, dto.Attempt{
			ID:     a.ID.String(),
			UserID: a.UserID.String(),
			ExamID: a.ExamID.String(),
			Status: status,
			Score:  a.TotalScore.Float64,
		})
	}

	return result, nil
}

// PrepareAttempt implements [AttemptService].
func (a *attemptService) PrepareAttempt(ctx context.Context, attemptID string) error {
	if attemptID == "" {
		return fmt.Errorf("attemptService.PrepareAttempt: attempt_id is required")
	}

	attemptUUID, err := uuid.Parse(attemptID)
	if err != nil {
		return fmt.Errorf("attemptService.PrepareAttempt: invalid attempt_id: %w", err)
	}

	// 1. get attempt to find exam_id
	attempt, err := a.attempt.GetAttempt(ctx, pgtype.UUID{Bytes: attemptUUID, Valid: true})
	if err != nil {
		return fmt.Errorf("attemptService.PrepareAttempt: get attempt: %w", err)
	}

	// 2. get all question sets for this exam
	questionSets, err := a.questionSet.GetQuestionSetsByExam(ctx, attempt.ExamID)
	if err != nil {
		return fmt.Errorf("attemptService.PrepareAttempt: get question sets: %w", err)
	}

	// 3. for each question set, fetch questions from Mongo
	var rows []db.BulkInsertAttemptMapParams
	orderIndex := 0

	for _, qs := range questionSets {
		questions, err := a.question.GetQuestionsBySet(ctx, qs.QuestionSetID)
		if err != nil {
			return fmt.Errorf("attemptService.PrepareAttempt: get questions for set %s: %w", qs.QuestionSetID, err)
		}

		// shuffle if needed — exam settings checked once outside loop ideally
		// for now shuffle at question level
		shuffleQuestions(questions)

		for _, q := range questions {
			// shuffle options per question if needed
			shuffledOpts, err := json.Marshal(shuffleOptions(q.Options))
			if err != nil {
				return fmt.Errorf("attemptService.PrepareAttempt: marshal shuffled options: %w", err)
			}

			sectionUUID, err := uuid.Parse(qs.SectionID.String())
			if err != nil {
				return fmt.Errorf("attemptService.PrepareAttempt: invalid section_id: %w", err)
			}

			rows = append(rows, db.BulkInsertAttemptMapParams{
				AttemptID:       pgtype.UUID{Bytes: attemptUUID, Valid: true},
				QuestionID:      q.ID,
				SectionID:       pgtype.UUID{Bytes: sectionUUID, Valid: true},
				OrderIndex:      pgtype.Int4{Int32: int32(orderIndex), Valid: true},
				ShuffledOptions: shuffledOpts,
			})
			orderIndex++
		}
	}

	// 4. bulk insert into attempt_question_map
	if _, err := a.attemptQuesMap.BulkInsertAttemptMap(ctx, rows); err != nil {
		return fmt.Errorf("attemptService.PrepareAttempt: bulk insert: %w", err)
	}

	return nil
}

// SaveAnswer implements [AttemptService].
func (a *attemptService) SaveAnswer(ctx context.Context, ans dto.Answer) error {
	if ans.AttemptID == "" {
		return fmt.Errorf("examService.SaveAnswer: attempt_id is required")
	}
	if ans.QuestionID == "" {
		return fmt.Errorf("examService.SaveAnswer: question_id is required")
	}
	if ans.SelectedAnswer == nil {
		return fmt.Errorf("examService.SaveAnswer: selected_answer is required")
	}

	attemptUUID, err := uuid.Parse(ans.AttemptID)
	if err != nil {
		return fmt.Errorf("examService.SaveAnswer: invalid attempt_id: %w", err)
	}

	questionUUID, err := uuid.Parse(ans.QuestionID)
	if err != nil {
		return fmt.Errorf("examService.SaveAnswer: invalid question_id: %w", err)
	}

	// marshal selected answer for JSONB storage
	selectedAnswer, err := json.Marshal(ans.SelectedAnswer)
	if err != nil {
		return fmt.Errorf("examService.SaveAnswer: marshal selected_answer: %w", err)
	}

	if err := a.answer.CreateAnswer(ctx, &db.UpsertAnswerParams{
		AttemptID:      pgtype.UUID{Bytes: attemptUUID, Valid: true},
		QuestionID:     questionUUID.String(),
		SelectedAnswer: selectedAnswer,
	}); err != nil {
		return fmt.Errorf("examService.SaveAnswer: %w", err)
	}

	return nil
}

// StartAttempt implements [AttemptService].
func (a *attemptService) StartAttempt(ctx context.Context, userID string, req dto.StartAttemptReq) (*dto.Attempt, error) {
	if userID == "" {
		return nil, fmt.Errorf("examService.StartAttempt: user_id is required")
	}
	if req.ExamID == "" {
		return nil, fmt.Errorf("examService.StartAttempt: exam_id is required")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("examService.StartAttempt: invalid user_id: %w", err)
	}

	examUUID, err := uuid.Parse(req.ExamID)
	if err != nil {
		return nil, fmt.Errorf("examService.StartAttempt: invalid exam_id: %w", err)
	}

	// check if there's already an in-progress attempt
	existing, err := a.attempt.GetActiveAttempt(ctx, &db.GetActiveAttemptParams{
		UserID: pgtype.UUID{Bytes: userUUID, Valid: true},
		ExamID: pgtype.UUID{Bytes: examUUID, Valid: true},
	})
	if err != nil && !isNotFound(err) {
		return nil, fmt.Errorf("examService.StartAttempt: check active attempt: %w", err)
	}

	if existing != nil {
		// return existing in-progress attempt instead of creating a new one
		status, ok := existing.Status.(string)
		if !ok {
			return nil, fmt.Errorf("attemptService.StartAttempt: unexpected Status type %T", existing.Status)
		}
		return &dto.Attempt{
			ID:     existing.ID.String(),
			UserID: existing.UserID.String(),
			ExamID: existing.ExamID.String(),
			Status: status,
			Score:  existing.TotalScore.Float64,
		}, nil
	}

	attempt, err := a.attempt.CreateAttempt(ctx, &db.CreateAttemptParams{
		UserID: pgtype.UUID{Bytes: userUUID, Valid: true},
		ExamID: pgtype.UUID{Bytes: examUUID, Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("examService.StartAttempt: %w", err)
	}

	// kick off question map preparation in background
	a.pool.Submit(func(ctx context.Context) error {
		return a.PrepareAttempt(ctx, attempt.ID.String())
	})

	status, ok := attempt.Status.(string)
	if !ok {
		return nil, fmt.Errorf("attemptService.StartAttempt: unexpected Status type %T", attempt.Status)
	}
	return &dto.Attempt{
		ID:     attempt.ID.String(),
		UserID: attempt.UserID.String(),
		ExamID: attempt.ExamID.String(),
		Status: status,
		Score:  attempt.TotalScore.Float64,
	}, nil
}

// SubmitAttempt implements [AttemptService].
func (a *attemptService) SubmitAttempt(ctx context.Context, attemptID string) (*dto.AttemptResult, error) {
	if attemptID == "" {
		return nil, fmt.Errorf("examService.SubmitAttempt: attempt_id is required")
	}

	attemptUUID, err := uuid.Parse(attemptID)
	if err != nil {
		return nil, fmt.Errorf("examService.SubmitAttempt: invalid attempt_id: %w", err)
	}

	// status check — only in_progress attempts can be submitted
	status, err := a.attempt.GetAttemptStatus(ctx, pgtype.UUID{Bytes: attemptUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("examService.SubmitAttempt: get status: %w", err)
	}
	if status.String != "in_progress" {
		return nil, fmt.Errorf("examService.SubmitAttempt: attempt is not in progress")
	}

	// kick off evaluation in background — scoring is heavy
	a.pool.Submit(func(ctx context.Context) error {
		return a.EvaluateAttempt(ctx, attemptID)
	})

	// return current (pre-score) result immediately
	return a.GetResult(ctx, attemptID)
}

func NewAttemptService(
	a *repository.AttemptRepository,
	aQM *repository.AttemptQuestionMapRepository,
	ans *repository.AnswerRepository,
	ques *repository.QuestionRepository,
	quesSet *repository.QuestionSetsRepository,
	score *repository.SectionScoresRepository,
	pool *worker.Pool,
) AttemptService {
	return &attemptService{
		attempt:        *a,
		attemptQuesMap: *aQM,
		answer:         *ans,
		question:       *ques,
		questionSet:    *quesSet,
		score:          *score,
		pool:           pool,
	}
}

func isNotFound(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

// hydrateQuestions fetches question content from Mongo and merges with PG map metadata
func (a *attemptService) hydrateQuestions(ctx context.Context, mapped []db.AttemptQuestionMap) ([]dto.QuestionWithMeta, error) {
	if len(mapped) == 0 {
		return []dto.QuestionWithMeta{}, nil
	}

	// collect question IDs
	ids := make([]string, 0, len(mapped))
	for _, m := range mapped {
		ids = append(ids, m.QuestionID)
	}

	// batch fetch from Mongo
	questions, err := a.question.GetQuestionsByIDs(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("hydrateQuestions: mongo fetch: %w", err)
	}

	// index by ID
	qMap := make(map[string]dto.Question, len(questions))
	for _, q := range questions {
		qMap[q.ID] = q
	}

	result := make([]dto.QuestionWithMeta, 0, len(mapped))
	for _, m := range mapped {
		q, ok := qMap[m.QuestionID]
		if !ok {
			continue // question deleted from Mongo — skip gracefully
		}

		// apply shuffled options if present
		if len(m.ShuffledOptions) > 0 {
			var opts []dto.Option
			if err := json.Unmarshal(m.ShuffledOptions, &opts); err == nil {
				q.Options = opts
			}
		}

		// strip correct answer — never send to client
		q.CorrectAnswer = nil

		result = append(result, dto.QuestionWithMeta{
			Question:  q,
			SectionID: m.SectionID.String(),
			Order:     int(m.OrderIndex.Int32),
		})
	}

	return result, nil
}

// isAnswerCorrect compares selected answer against correct answer by question type
func isAnswerCorrect(q dto.Question, selected interface{}) bool {
	if selected == nil {
		return false
	}
	switch q.Type {
	case "MCQ":
		return fmt.Sprintf("%v", selected) == fmt.Sprintf("%v", q.CorrectAnswer)
	case "MSQ":
		// both should be []interface{} — compare as sets
		sel, ok1 := toStringSet(selected)
		cor, ok2 := toStringSet(q.CorrectAnswer)
		if !ok1 || !ok2 {
			return false
		}
		return equalSets(sel, cor)
	case "NUMERIC":
		return fmt.Sprintf("%v", selected) == fmt.Sprintf("%v", q.CorrectAnswer)
	default:
		return false
	}
}

func calculateMarks(q dto.Question, isCorrect bool) float64 {
	if isCorrect {
		return q.Marks
	}
	return -q.NegativeMarks // already 0 if no negative marking
}

func toStringSet(v interface{}) (map[string]struct{}, bool) {
	arr, ok := v.([]interface{})
	if !ok {
		return nil, false
	}
	set := make(map[string]struct{}, len(arr))
	for _, item := range arr {
		set[fmt.Sprintf("%v", item)] = struct{}{}
	}
	return set, true
}

func equalSets(a, b map[string]struct{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}

func shuffleQuestions(questions []dto.Question) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}

func shuffleOptions(options []dto.Option) []dto.Option {
	shuffled := make([]dto.Option, len(options))
	copy(shuffled, options)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
