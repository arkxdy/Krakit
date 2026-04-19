package dto

type AttachQuestionSetReq struct {
	ExamID        string         `json:"exam_id"`
	SectionID     string         `json:"section_id"`
	QuestionSetID string         `json:"question_set_id"`
	QuestionCount int            `json:"question_count"`
	Difficulty    map[string]int `json:"difficulty_distribution"`
}
