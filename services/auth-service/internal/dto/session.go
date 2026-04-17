package dto

type SessionResponse struct {
	ID        string `json:"id"`
	Platform  string `json:"platform"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at"`
	IsCurrent bool   `json:"is_current"`
}

type ListSessionsResponse struct {
	Sessions []SessionResponse `json:"sessions"`
}
