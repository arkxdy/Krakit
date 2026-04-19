package dto

type CreateExamReq struct {
	Name            string `json:"name"`
	ExamType        string `json:"exam_type"`
	DurationMinutes int    `json:"duration_minutes"`
	TotalMarks      int    `json:"total_marks"`
}

type UpdateExamReq struct {
	ExamID          string `json:"exam_id"`
	Name            string `json:"name"`
	DurationMinutes int    `json:"duration_minutes"`
	TotalMarks      int    `json:"total_marks"`
}

type Exam struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ExamType        string `json:"exam_type"`
	DurationMinutes int    `json:"duration_minutes"`
	TotalMarks      int    `json:"total_marks"`
	IsActive        bool   `json:"is_active"`
}
