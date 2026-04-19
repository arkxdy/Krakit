package dto

type Option struct {
	Key   string `json:"key"`
	Text  string `json:"text"`
	Image string `json:"image,omitempty"`
}
