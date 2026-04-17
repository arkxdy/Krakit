package dto

// ---------- REQUESTS ----------

type SignupRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Platform  string `json:"platform" binding:"required,oneof=web mobile ios"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Platform string `json:"platform" binding:"required,oneof=web mobile ios"`
}

// refresh comes from cookie → no body needed
type RefreshRequest struct{}

type LogoutRequest struct{}

type LogoutAllRequest struct{}

// ---------- RESPONSES ----------

type AuthUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Plan     string `json:"plan"`
}

type AuthResponse struct {
	User        AuthUser `json:"user"`
	AccessToken string   `json:"access_token"`
	ExpiresIn   int64    `json:"expires_in"`
}

type GoogleAuthRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}
