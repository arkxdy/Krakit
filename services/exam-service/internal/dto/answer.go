package dto

type Answer struct {
	AttemptID      string      `json:"attempt_id"`
	QuestionID     string      `json:"question_id"`
	SelectedAnswer interface{} `json:"selected_answer"`
}
