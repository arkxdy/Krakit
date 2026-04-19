package middleware

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// =========================
// JWKS CLIENT
// =========================

type jwk struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type jwksResponse struct {
	Keys []jwk `json:"keys"`
}

type JWKSClient struct {
	url        string
	mu         sync.RWMutex
	keys       map[string]*rsa.PublicKey
	lastFetch  time.Time
	cacheTTL   time.Duration
	httpClient *http.Client
}

func NewJWKSClient(jwksURL string) *JWKSClient {
	return &JWKSClient{
		url:        jwksURL,
		keys:       make(map[string]*rsa.PublicKey),
		cacheTTL:   10 * time.Minute,
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

func (j *JWKSClient) getKey(kid string) (*rsa.PublicKey, error) {
	// check cache first
	j.mu.RLock()
	if time.Since(j.lastFetch) < j.cacheTTL {
		if key, ok := j.keys[kid]; ok {
			j.mu.RUnlock()
			return key, nil
		}
	}
	j.mu.RUnlock()

	// cache miss or expired — refetch
	if err := j.fetchKeys(); err != nil {
		return nil, fmt.Errorf("jwks: fetch keys: %w", err)
	}

	j.mu.RLock()
	defer j.mu.RUnlock()

	key, ok := j.keys[kid]
	if !ok {
		return nil, fmt.Errorf("jwks: key %q not found", kid)
	}

	return key, nil
}

func (j *JWKSClient) fetchKeys() error {
	resp, err := j.httpClient.Get(j.url)
	if err != nil {
		return fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	var jwks jwksResponse
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	keys := make(map[string]*rsa.PublicKey, len(jwks.Keys))
	for _, k := range jwks.Keys {
		if k.Kty != "RSA" || k.Use != "sig" {
			continue
		}

		pub, err := parseRSAPublicKey(k.N, k.E)
		if err != nil {
			return fmt.Errorf("parse key %s: %w", k.Kid, err)
		}

		keys[k.Kid] = pub
	}

	j.mu.Lock()
	j.keys = keys
	j.lastFetch = time.Now()
	j.mu.Unlock()

	return nil
}

func parseRSAPublicKey(nStr, eStr string) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(nStr)
	if err != nil {
		return nil, fmt.Errorf("decode n: %w", err)
	}

	eBytes, err := base64.RawURLEncoding.DecodeString(eStr)
	if err != nil {
		return nil, fmt.Errorf("decode e: %w", err)
	}

	n := new(big.Int).SetBytes(nBytes)
	e := new(big.Int).SetBytes(eBytes)

	return &rsa.PublicKey{
		N: n,
		E: int(e.Int64()),
	}, nil
}

// =========================
// JWT CLAIMS
// same shape as auth-service JWTClaims
// =========================

type Claims struct {
	UserID      string   `json:"user_id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

// =========================
// MIDDLEWARE
// =========================

func AuthMiddleware(jwks *JWKSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := extractBearer(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := verifyToken(tokenStr, jwks)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

func verifyToken(tokenStr string, jwks *JWKSClient) (*Claims, error) {
	// parse header to get kid before full verification
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing kid in token header")
		}

		return jwks.getKey(kid)
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}

func extractBearer(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization format")
	}

	return parts[1], nil
}

// =========================
// HELPERS — same pattern as auth-service
// =========================

func GetUserFromContext(c *gin.Context) (*Claims, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	claims, ok := user.(*Claims)
	return claims, ok
}

func RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := GetUserFromContext(c)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		for _, p := range user.Permissions {
			if p == permission {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "permission denied"})
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := GetUserFromContext(c)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		if user.Role != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not allowed"})
			return
		}
		c.Next()
	}
}

func RequireAnyRole(roles ...string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(roles))
	for _, r := range roles {
		allowed[r] = struct{}{}
	}
	return func(c *gin.Context) {
		user, exists := GetUserFromContext(c)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		if _, ok := allowed[user.Role]; !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not allowed"})
			return
		}
		c.Next()
	}
}

func RequireAdmin() gin.HandlerFunc {
	return RequireAnyRole("admin", "super_admin")
}
