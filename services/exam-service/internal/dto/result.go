package dto

type SectionResult struct {
	SectionID string  `json:"section_id"`
	Score     float64 `json:"score"`
	Correct   int     `json:"correct"`
	Wrong     int     `json:"wrong"`
}

type AttemptResult struct {
	AttemptID string          `json:"attempt_id"`
	Score     float64         `json:"score"`
	Sections  []SectionResult `json:"sections"`
}
