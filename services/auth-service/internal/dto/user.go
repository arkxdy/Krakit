package dto

// ---------- REQUESTS ----------
type GetUserRequest struct{}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// ---------- RESPONSES ----------

type UserResponse struct {
	ID          string  `json:"id"`
	Email       string  `json:"email"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	FullName    string  `json:"full_name"`
	Role        string  `json:"role"`
	Plan        string  `json:"plan"`
	LastLoginAt *string `json:"last_login_at,omitempty"`
}
