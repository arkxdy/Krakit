package dto

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message,omitempty"`
}

type ErrorCode string

func Success(data interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Data:    data,
		Message: message,
	}
}

func Error(code ErrorCode, message string) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}

const (
	// 🔹 General
	ErrInvalidRequest ErrorCode = "INVALID_REQUEST"
	ErrBadRequest     ErrorCode = "BAD_REQUEST"
	ErrInternal       ErrorCode = "INTERNAL_ERROR"
	ErrNotFound       ErrorCode = "NOT_FOUND"
	ErrConflict       ErrorCode = "CONFLICT"

	// 🔹 Auth
	ErrUnauthorized   ErrorCode = "UNAUTHORIZED"
	ErrForbidden      ErrorCode = "FORBIDDEN"
	ErrInvalidCreds   ErrorCode = "INVALID_CREDENTIALS"
	ErrTokenExpired   ErrorCode = "TOKEN_EXPIRED"
	ErrInvalidToken   ErrorCode = "INVALID_TOKEN"
	ErrSessionExpired ErrorCode = "SESSION_EXPIRED"
	ErrSessionRevoked ErrorCode = "SESSION_REVOKED"

	// 🔹 User
	ErrUserNotFound     ErrorCode = "USER_NOT_FOUND"
	ErrUserAlreadyExist ErrorCode = "USER_ALREADY_EXISTS"

	// 🔹 Validation
	ErrValidationFailed ErrorCode = "VALIDATION_FAILED"
	ErrMissingField     ErrorCode = "MISSING_FIELD"
	ErrInvalidField     ErrorCode = "INVALID_FIELD"

	// 🔹 Permissions / RBAC
	ErrPermissionDenied ErrorCode = "PERMISSION_DENIED"
	ErrRoleNotAllowed   ErrorCode = "ROLE_NOT_ALLOWED"

	// 🔹 Rate limiting / security
	ErrTooManyRequests ErrorCode = "TOO_MANY_REQUESTS"
	ErrTokenReuse      ErrorCode = "TOKEN_REUSE_DETECTED"

	// 🔹 External services
	ErrGoogleAuthFailed   ErrorCode = "GOOGLE_AUTH_FAILED"
	ErrServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
)
