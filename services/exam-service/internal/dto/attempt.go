package dto

type Attempt struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	ExamID string  `json:"exam_id"`
	Status string  `json:"status"`
	Score  float64 `json:"score"`
}

type StartAttemptReq struct {
	ExamID string `json:"exam_id"`
}
