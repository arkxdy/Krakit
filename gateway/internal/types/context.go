package types

type ContextKey string

const (
	UserIDKey    ContextKey = "user_id"
	RoleKey      ContextKey = "role"
	RequestIDKey ContextKey = "request_id"
)
