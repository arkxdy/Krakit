package dto

type Passage struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	ExamID      string `json:"exam_id"`
	SectionID   string `json:"section_id"`
	SubjectID   string `json:"subject_id,omitempty"`
	PassageText string `json:"passage_text"`
	Topic       string `json:"topic,omitempty"`
	IsActive    bool   `json:"is_active"`
}
