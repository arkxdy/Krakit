package repository

import (
	"context"
	"fmt"

	"github.com/krakit/exam-service/internal/db"
	"github.com/krakit/exam-service/internal/dto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	dbName              = "krakit"
	questionsCollection = "questions"
	passagesCollection  = "passages"
)

type questionRepo struct {
	questions *mongo.Collection
	passages  *mongo.Collection
}

// BulkUpsertPassages implements [QuestionRepository].
func (q *questionRepo) BulkUpsertPassages(ctx context.Context, passages []dto.Passage) error {
	if len(passages) == 0 {
		return nil
	}

	models := make([]mongo.WriteModel, 0, len(passages))
	for _, p := range passages {
		filter := bson.M{"_id": p.ID}
		update := bson.M{"$set": p}
		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)
		models = append(models, model)
	}

	opts := options.BulkWrite().SetOrdered(false)
	if _, err := q.passages.BulkWrite(ctx, models, opts); err != nil {
		return fmt.Errorf("questionRepo.BulkUpsertPassages: %w", err)
	}

	return nil
}

// BulkUpsertQuestions implements [QuestionRepository].
func (q *questionRepo) BulkUpsertQuestions(ctx context.Context, questions []dto.Question) error {
	if len(questions) == 0 {
		return nil
	}

	models := make([]mongo.WriteModel, 0, len(questions))
	for _, q := range questions {
		filter := bson.M{"_id": q.ID}
		update := bson.M{"$set": q}
		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(true)
		models = append(models, model)
	}

	opts := options.BulkWrite().SetOrdered(false) // unordered = faster, all ops attempted
	if _, err := q.questions.BulkWrite(ctx, models, opts); err != nil {
		return fmt.Errorf("questionRepo.BulkUpsertQuestions: %w", err)
	}

	return nil
}

// DeletePassagesBySection implements [QuestionRepository].
func (q *questionRepo) DeletePassagesBySection(ctx context.Context, sectionID string) error {
	filter := bson.M{"section_id": sectionID}
	update := bson.M{"$set": bson.M{"is_active": false}}

	if _, err := q.passages.UpdateMany(ctx, filter, update); err != nil {
		return fmt.Errorf("questionRepo.DeletePassagesBySection: %w", err)
	}

	return nil
}

// DeleteQuestionsBySet implements [QuestionRepository].
func (q *questionRepo) DeleteQuestionsBySet(ctx context.Context, setID string) error {
	filter := bson.M{"question_set_id": setID}
	update := bson.M{"$set": bson.M{"is_active": false}}

	if _, err := q.questions.UpdateMany(ctx, filter, update); err != nil {
		return fmt.Errorf("questionRepo.DeleteQuestionsBySet: %w", err)
	}

	return nil
}

// GetPassagesByIDs implements [QuestionRepository].
func (q *questionRepo) GetPassagesByIDs(ctx context.Context, ids []string) ([]dto.Passage, error) {
	if len(ids) == 0 {
		return []dto.Passage{}, nil
	}

	filter := bson.M{
		"_id": bson.M{"$in": ids},
	}

	cursor, err := q.passages.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("questionRepo.GetPassagesByIDs: %w", err)
	}
	defer cursor.Close(ctx)

	var passages []dto.Passage
	if err := cursor.All(ctx, &passages); err != nil {
		return nil, fmt.Errorf("questionRepo.GetPassagesByIDs: decode: %w", err)
	}

	return passages, nil
}

// GetQuestionsByIDs implements [QuestionRepository].
func (q *questionRepo) GetQuestionsByIDs(ctx context.Context, ids []string) ([]dto.Question, error) {
	if len(ids) == 0 {
		return []dto.Question{}, nil
	}

	filter := bson.M{
		"_id": bson.M{"$in": ids},
	}

	cursor, err := q.questions.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("questionRepo.GetQuestionsByIDs: %w", err)
	}
	defer cursor.Close(ctx)

	var questions []dto.Question
	if err := cursor.All(ctx, &questions); err != nil {
		return nil, fmt.Errorf("questionRepo.GetQuestionsByIDs: decode: %w", err)
	}

	return questions, nil
}

// GetQuestionsBySet implements [QuestionRepository].
func (q *questionRepo) GetQuestionsBySet(ctx context.Context, setID string) ([]dto.Question, error) {
	filter := bson.M{
		"question_set_id": setID,
		"is_active":       true,
	}

	cursor, err := q.questions.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("questionRepo.GetQuestionsBySet: %w", err)
	}
	defer cursor.Close(ctx)

	var questions []dto.Question
	if err := cursor.All(ctx, &questions); err != nil {
		return nil, fmt.Errorf("questionRepo.GetQuestionsBySet: decode: %w", err)
	}

	return questions, nil
}

// InsertPassages implements [QuestionRepository].
func (q *questionRepo) InsertPassages(ctx context.Context, passages []dto.Passage) error {
	if len(passages) == 0 {
		return nil
	}

	docs := make([]interface{}, 0, len(passages))
	for _, p := range passages {
		docs = append(docs, p)
	}

	if _, err := q.passages.InsertMany(ctx, docs); err != nil {
		return fmt.Errorf("questionRepo.InsertPassages: %w", err)
	}

	return nil
}

// InsertQuestions implements [QuestionRepository].
func (q *questionRepo) InsertQuestions(ctx context.Context, questions []dto.Question) error {
	if len(questions) == 0 {
		return nil
	}

	docs := make([]interface{}, 0, len(questions))
	for _, q := range questions {
		docs = append(docs, q)
	}

	if _, err := q.questions.InsertMany(ctx, docs); err != nil {
		return fmt.Errorf("questionRepo.InsertQuestions: %w", err)
	}

	return nil
}

type QuestionRepository interface {

	// =========================
	// QUESTIONS
	// =========================

	InsertQuestions(ctx context.Context, questions []dto.Question) error
	BulkUpsertQuestions(ctx context.Context, questions []dto.Question) error

	GetQuestionsBySet(ctx context.Context, setID string) ([]dto.Question, error)
	GetQuestionsByIDs(ctx context.Context, ids []string) ([]dto.Question, error)

	DeleteQuestionsBySet(ctx context.Context, setID string) error

	// =========================
	// PASSAGES
	// =========================

	InsertPassages(ctx context.Context, passages []dto.Passage) error
	BulkUpsertPassages(ctx context.Context, passages []dto.Passage) error

	GetPassagesByIDs(ctx context.Context, ids []string) ([]dto.Passage, error)

	DeletePassagesBySection(ctx context.Context, sectionID string) error
}

func NewQuestionRepository(mongo *db.MongoClient) QuestionRepository {
	return &questionRepo{
		questions: mongo.Client.Database(dbName).Collection(questionsCollection),
		passages:  mongo.Client.Database(dbName).Collection(passagesCollection),
	}
}
