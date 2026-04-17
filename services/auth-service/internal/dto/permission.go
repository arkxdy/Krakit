package dto

type PermissionResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AssignPermissionRequest struct {
	PermissionID string `json:"permission_id" binding:"required,uuid"`
}
