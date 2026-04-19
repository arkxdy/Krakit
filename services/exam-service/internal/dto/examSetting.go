package dto

type ExamSettings struct {
	ExamID             string `json:"exam_id"`
	IsPublished        bool   `json:"is_published"`
	ShuffleQuestions   bool   `json:"shuffle_questions"`
	ShuffleOptions     bool   `json:"shuffle_options"`
	AllowSectionSwitch bool   `json:"allow_section_switch"`
}
