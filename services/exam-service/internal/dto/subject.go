package dto

type CreateSubjectReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Subject struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
