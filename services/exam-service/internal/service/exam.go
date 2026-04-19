package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/exam-service/internal/db/sqlc"
	"github.com/krakit/exam-service/internal/dto"
	"github.com/krakit/exam-service/internal/repository"
)

type ExamService interface {

	// =====================================================
	// ADMIN APIs
	// =====================================================

	// SUBJECTS
	CreateSubject(ctx context.Context, req dto.CreateSubjectReq) (*dto.Subject, error)
	GetSubjects(ctx context.Context) ([]dto.Subject, error)

	// EXAMS
	CreateExam(ctx context.Context, req dto.CreateExamReq) (*dto.Exam, error)
	UpdateExam(ctx context.Context, req dto.UpdateExamReq) error
	ListExams(ctx context.Context) ([]dto.Exam, error)
	ListExamsPaginated(ctx context.Context, limit, offset int32) ([]dto.Exam, error)
	DisableExam(ctx context.Context, examID string) error

	// SETTINGS
	CreateExamSettings(ctx context.Context, settings dto.ExamSettings) error
	PublishExam(ctx context.Context, examID string) error

	// SECTIONS
	CreateSection(ctx context.Context, req dto.CreateSectionReq) (*dto.Section, error)
	UpdateSection(ctx context.Context, sectionID string, req dto.CreateSectionReq) error
	GetSections(ctx context.Context, examID string) ([]dto.Section, error)
}

type examService struct {
	exam    repository.ExamRepository
	setting repository.ExamSettingsRepository
	section repository.ExamSectionRepository
	subject repository.SubjectRepository
}

// CreateExam implements [ExamService].
func (e *examService) CreateExam(ctx context.Context, req dto.CreateExamReq) (*dto.Exam, error) {
	// validate
	if req.Name == "" {
		return nil, fmt.Errorf("examService.CreateExam: name is required")
	}
	if req.ExamType == "" {
		return nil, fmt.Errorf("examService.CreateExam: exam type is required")
	}
	if req.DurationMinutes <= 0 {
		return nil, fmt.Errorf("examService.CreateExam: duration must be positive")
	}
	if req.TotalMarks <= 0 {
		return nil, fmt.Errorf("examService.CreateExam: total marks must be positive")
	}

	args := db.CreateExamParams{
		Name:            req.Name,
		ExamType:        req.ExamType,
		DurationMinutes: int32(req.DurationMinutes),
		TotalMarks:      int32(req.TotalMarks),
	}

	exam, err := e.exam.CreateExam(ctx, &args)
	if err != nil {
		return nil, fmt.Errorf("examService.CreateExam: %w", err)
	}

	// ExamType is a string column — sqlc generates it as string directly
	// if your sqlc generates a custom type, change this accordingly
	examType, ok := exam.ExamType.(string)
	if !ok {
		return nil, fmt.Errorf("examService.CreateExam: unexpected ExamType type %T", exam.ExamType)
	}

	return &dto.Exam{
		ID:              exam.ID.String(),
		Name:            exam.Name,
		ExamType:        examType,
		DurationMinutes: int(exam.DurationMinutes),
		TotalMarks:      int(exam.TotalMarks),
		IsActive:        exam.IsActive.Bool,
	}, nil
}

// CreateExamSettings implements [ExamService].
func (e *examService) CreateExamSettings(ctx context.Context, settings dto.ExamSettings) error {
	// validate
	if settings.ExamID == "" {
		return fmt.Errorf("examService.CreateExamSettings: exam_id is required")
	}

	examUUID, err := uuid.Parse(settings.ExamID)
	if err != nil {
		return fmt.Errorf("examService.CreateExamSettings: invalid exam_id: %w", err)
	}

	args := db.UpsertExamSettingsParams{
		ExamID:             pgtype.UUID{Bytes: examUUID, Valid: true},
		IsPublished:        pgtype.Bool{Bool: settings.IsPublished, Valid: true},
		ShuffleQuestions:   pgtype.Bool{Bool: settings.ShuffleQuestions, Valid: true},
		ShuffleOptions:     pgtype.Bool{Bool: settings.ShuffleOptions, Valid: true},
		AllowSectionSwitch: pgtype.Bool{Bool: settings.AllowSectionSwitch, Valid: true},
	}

	_, err = e.setting.CreateExamSettings(ctx, &args)
	if err != nil {
		return fmt.Errorf("examService.CreateExamSettings: %w", err)
	}

	return nil
}

// CreateSection implements [ExamService].
func (e *examService) CreateSection(ctx context.Context, req dto.CreateSectionReq) (*dto.Section, error) {
	// validate
	if req.ExamID == "" {
		return nil, fmt.Errorf("examService.CreateSection: exam_id is required")
	}
	if req.SubjectID == "" {
		return nil, fmt.Errorf("examService.CreateSection: subject_id is required")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("examService.CreateSection: name is required")
	}
	if req.QuestionCount <= 0 {
		return nil, fmt.Errorf("examService.CreateSection: question_count must be positive")
	}
	if req.OrderIndex < 0 {
		return nil, fmt.Errorf("examService.CreateSection: order_index must be non-negative")
	}

	examUUID, err := uuid.Parse(req.ExamID)
	if err != nil {
		return nil, fmt.Errorf("examService.CreateSection: invalid exam_id: %w", err)
	}

	subjectUUID, err := uuid.Parse(req.SubjectID)
	if err != nil {
		return nil, fmt.Errorf("examService.CreateSection: invalid subject_id: %w", err)
	}

	args := db.CreateSectionParams{
		ExamID:          pgtype.UUID{Bytes: examUUID, Valid: true},
		Name:            req.Name,
		SubjectID:       pgtype.UUID{Bytes: subjectUUID, Valid: true},
		TimeLimit:       pgtype.Int4{Int32: int32(req.TimeLimit), Valid: true},
		QuestionCount:   pgtype.Int4{Int32: int32(req.QuestionCount), Valid: true},
		OrderIndex:      pgtype.Int4{Int32: int32(req.OrderIndex), Valid: true},
		IsSwitchAllowed: pgtype.Bool{Bool: req.IsSwitchAllowed, Valid: true},
	}

	section, err := e.section.CreateSection(ctx, &args)
	if err != nil {
		return nil, fmt.Errorf("examService.CreateSection: %w", err)
	}

	return &dto.Section{
		ID:              section.ID.String(),
		Name:            section.Name,
		ExamID:          section.ExamID.String(),
		SubjectID:       section.SubjectID.String(),
		TimeLimit:       int(section.TimeLimit.Int32),
		QuestionCount:   int(section.QuestionCount.Int32),
		OrderIndex:      int(section.OrderIndex.Int32),
		IsSwitchAllowed: section.IsSwitchAllowed.Bool,
	}, nil
}

// CreateSubject implements [ExamService].
func (e *examService) CreateSubject(ctx context.Context, req dto.CreateSubjectReq) (*dto.Subject, error) {
	//validate
	if req.Name == "" {
		return nil, fmt.Errorf("examService.CreateSubject: name is required")
	}

	args := db.CreateSubjectParams{
		Name:        req.Name,
		Description: pgtype.Text{String: req.Description, Valid: req.Description != ""},
	}

	subject, err := e.subject.CreateSubject(ctx, &args)
	if err != nil {
		return nil, fmt.Errorf("examService.CreateSubject: %w", err)
	}

	return &dto.Subject{
		ID:          subject.ID.String(),
		Name:        subject.Name,
		Description: subject.Description.String,
	}, nil
}

// DisableExam implements [ExamService].
func (e *examService) DisableExam(ctx context.Context, examID string) error {
	if examID == "" {
		return fmt.Errorf("examService.DisableExam: exam_id is required")
	}

	examUUID, err := uuid.Parse(examID)
	if err != nil {
		return fmt.Errorf("examService.DisableExam: invalid exam_id: %w", err)
	}

	if err := e.exam.DisableExam(ctx, pgtype.UUID{Bytes: examUUID, Valid: true}); err != nil {
		return fmt.Errorf("examService.DisableExam: %w", err)
	}

	return nil
}

// GetSections implements [ExamService].
func (e *examService) GetSections(ctx context.Context, examID string) ([]dto.Section, error) {
	if examID == "" {
		return nil, fmt.Errorf("examService.GetSections: exam_id is required")
	}

	examUUID, err := uuid.Parse(examID)
	if err != nil {
		return nil, fmt.Errorf("examService.GetSections: invalid exam_id: %w", err)
	}

	sections, err := e.section.GetSectionsByExam(ctx, pgtype.UUID{Bytes: examUUID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("examService.GetSections: %w", err)
	}

	result := make([]dto.Section, 0, len(sections))
	for _, s := range sections {
		result = append(result, dto.Section{
			ID:              s.ID.String(),
			Name:            s.Name,
			ExamID:          s.ExamID.String(),
			SubjectID:       s.SubjectID.String(),
			TimeLimit:       int(s.TimeLimit.Int32),
			QuestionCount:   int(s.QuestionCount.Int32),
			OrderIndex:      int(s.OrderIndex.Int32),
			IsSwitchAllowed: s.IsSwitchAllowed.Bool,
		})
	}

	return result, nil
}

// GetSubjects implements [ExamService].
func (e *examService) GetSubjects(ctx context.Context) ([]dto.Subject, error) {
	subjects, err := e.subject.GetSubjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("examService.GetSubjects: %w", err)
	}

	result := make([]dto.Subject, 0, len(subjects))
	for _, s := range subjects {
		result = append(result, dto.Subject{
			ID:          s.ID.String(),
			Name:        s.Name,
			Description: s.Description.String,
		})
	}

	return result, nil
}

// ListExams implements [ExamService].
func (e *examService) ListExams(ctx context.Context) ([]dto.Exam, error) {
	exams, err := e.exam.ListExams(ctx)
	if err != nil {
		return nil, fmt.Errorf("examService.ListExams: %w", err)
	}

	result := make([]dto.Exam, 0, len(exams))
	for _, ex := range exams {
		examType, ok := ex.ExamType.(string)
		if !ok {
			return nil, fmt.Errorf("examService.ListExams: unexpected ExamType type %T", ex.ExamType)
		}
		result = append(result, dto.Exam{
			ID:              ex.ID.String(),
			Name:            ex.Name,
			ExamType:        examType,
			DurationMinutes: int(ex.DurationMinutes),
			TotalMarks:      int(ex.TotalMarks),
			IsActive:        ex.IsActive.Bool,
		})
	}

	return result, nil
}

// ListExamsPaginated implements [ExamService].
func (e *examService) ListExamsPaginated(ctx context.Context, limit int32, offset int32) ([]dto.Exam, error) {
	if limit <= 0 {
		return nil, fmt.Errorf("examService.ListExamsPaginated: limit must be positive")
	}
	if offset < 0 {
		return nil, fmt.Errorf("examService.ListExamsPaginated: offset must be non-negative")
	}

	exams, err := e.exam.ListExamsPaginated(ctx, &db.ListExamsPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("examService.ListExamsPaginated: %w", err)
	}

	result := make([]dto.Exam, 0, len(exams))
	for _, ex := range exams {
		examType, ok := ex.ExamType.(string)
		if !ok {
			return nil, fmt.Errorf("examService.ListExamsPaginated: unexpected ExamType type %T", ex.ExamType)
		}
		result = append(result, dto.Exam{
			ID:              ex.ID.String(),
			Name:            ex.Name,
			ExamType:        examType,
			DurationMinutes: int(ex.DurationMinutes),
			TotalMarks:      int(ex.TotalMarks),
			IsActive:        ex.IsActive.Bool,
		})
	}

	return result, nil
}

// PublishExam implements [ExamService].
func (e *examService) PublishExam(ctx context.Context, examID string) error {
	if examID == "" {
		return fmt.Errorf("examService.PublishExam: exam_id is required")
	}

	examUUID, err := uuid.Parse(examID)
	if err != nil {
		return fmt.Errorf("examService.PublishExam: invalid exam_id: %w", err)
	}

	if err := e.setting.PublishExam(ctx, pgtype.UUID{Bytes: examUUID, Valid: true}); err != nil {
		return fmt.Errorf("examService.PublishExam: %w", err)
	}

	return nil
}

// UpdateExam implements [ExamService].
func (e *examService) UpdateExam(ctx context.Context, req dto.UpdateExamReq) error {
	if req.ExamID == "" {
		return fmt.Errorf("examService.UpdateExam: id is required")
	}
	if req.Name == "" {
		return fmt.Errorf("examService.UpdateExam: name is required")
	}
	if req.DurationMinutes <= 0 {
		return fmt.Errorf("examService.UpdateExam: duration must be positive")
	}
	if req.TotalMarks <= 0 {
		return fmt.Errorf("examService.UpdateExam: total marks must be positive")
	}

	examUUID, err := uuid.Parse(req.ExamID)
	if err != nil {
		return fmt.Errorf("examService.UpdateExam: invalid id: %w", err)
	}

	if err := e.exam.UpdateExam(ctx, &db.UpdateExamParams{
		ID:              pgtype.UUID{Bytes: examUUID, Valid: true},
		Name:            req.Name,
		DurationMinutes: int32(req.DurationMinutes),
		TotalMarks:      int32(req.TotalMarks),
	}); err != nil {
		return fmt.Errorf("examService.UpdateExam: %w", err)
	}

	return nil
}

// UpdateSection implements [ExamService].
func (e *examService) UpdateSection(ctx context.Context, sectionID string, req dto.CreateSectionReq) error {
	if sectionID == "" {
		return fmt.Errorf("examService.UpdateSection: section_id is required")
	}
	if req.Name == "" {
		return fmt.Errorf("examService.UpdateSection: name is required")
	}
	if req.QuestionCount <= 0 {
		return fmt.Errorf("examService.UpdateSection: question_count must be positive")
	}
	if req.OrderIndex < 0 {
		return fmt.Errorf("examService.UpdateSection: order_index must be non-negative")
	}

	sectionUUID, err := uuid.Parse(sectionID)
	if err != nil {
		return fmt.Errorf("examService.UpdateSection: invalid section_id: %w", err)
	}

	if err := e.section.UpdateSection(ctx, &db.UpdateSectionParams{
		ID:            pgtype.UUID{Bytes: sectionUUID, Valid: true},
		Name:          req.Name,
		TimeLimit:     pgtype.Int4{Int32: int32(req.TimeLimit), Valid: true},
		QuestionCount: pgtype.Int4{Int32: int32(req.QuestionCount), Valid: true},
		OrderIndex:    pgtype.Int4{Int32: int32(req.OrderIndex), Valid: true},
	}); err != nil {
		return fmt.Errorf("examService.UpdateSection: %w", err)
	}

	return nil
}

func NewExamService(
	exam *repository.ExamRepository,
	setting *repository.ExamSettingsRepository,
	section *repository.ExamSectionRepository,
	subject *repository.SubjectRepository,
) ExamService {
	return &examService{
		exam:    *exam,
		setting: *setting,
		section: *section,
		subject: *subject,
	}
}
