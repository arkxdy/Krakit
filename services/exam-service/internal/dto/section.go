package dto

type CreateSectionReq struct {
	ExamID          string `json:"exam_id"`
	Name            string `json:"name"`
	SubjectID       string `json:"subject_id"`
	TimeLimit       int    `json:"time_limit"`
	QuestionCount   int    `json:"question_count"`
	OrderIndex      int    `json:"order_index"`
	IsSwitchAllowed bool   `json:"is_switch_allowed"`
}

type Section struct {
	ID              string `json:"id"`
	ExamID          string `json:"exam_id"`
	Name            string `json:"name"`
	SubjectID       string `json:"subject_id"`
	TimeLimit       int    `json:"time_limit"`
	QuestionCount   int    `json:"question_count"`
	OrderIndex      int    `json:"order_index"`
	IsSwitchAllowed bool   `json:"is_switch_allowed"`
}
