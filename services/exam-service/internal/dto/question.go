package dto

type Question struct {
	ID            string `json:"id" bson:"_id,omitempty"`
	ExamID        string `json:"exam_id"`
	SectionID     string `json:"section_id"`
	SubjectID     string `json:"subject_id,omitempty"`
	QuestionSetID string `json:"question_set_id"`

	PassageID *string `json:"passage_id,omitempty"`

	QuestionText  string `json:"question_text"`
	QuestionImage string `json:"question_image_url,omitempty"`

	Options []Option `json:"options"`

	CorrectAnswer interface{} `json:"correct_answer"`

	Type string `json:"type"` // MCQ, MSQ, NUMERIC

	Marks         float64 `json:"marks"`
	NegativeMarks float64 `json:"negative_marks"`

	Difficulty string `json:"difficulty"`

	Topic    string `json:"topic,omitempty"`
	IsActive bool   `json:"is_active"`
}

type QuestionWithMeta struct {
	Question
	SectionID string `json:"section_id"`
	Order     int    `json:"order"`
}
