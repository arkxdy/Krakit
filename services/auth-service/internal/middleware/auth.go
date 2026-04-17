package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/utils"
)

func AuthMiddleware(jwtMaker *utils.JWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr, err := extractBearer(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		claims, err := jwtMaker.VerifyToken(tokenStr)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", claims)
		c.Next()
	}
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

func GetUserFromContext(c *gin.Context) (*dto.JWTClaims, bool) {
	user, exists := c.Get(utils.ContextUserKey)
	if !exists {
		return nil, false
	}

	claims, ok := user.(*dto.JWTClaims)
	return claims, ok
}
