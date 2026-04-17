package dto

type JWTClaims struct {
	Sub         string   `json:"sub"`
	Email       string   `json:"email"`
	Role        string   `json:"role"`
	Plan        string   `json:"plan"`
	Permissions []string `json:"permissions"`
	JTI         string   `json:"jti"`
	IssuedAt    int64    `json:"iat"`
	ExpiresAt   int64    `json:"exp"`
}

type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type JWKSResponse struct {
	Keys []JWK `json:"keys"`
}
