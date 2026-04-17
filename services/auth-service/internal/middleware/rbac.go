package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakit/auth-service/internal/dto"
)

func RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := GetUserFromContext(c)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error(dto.ErrUnauthorized, "User not authenticated"))
			return
		}

		for _, p := range user.Permissions {
			if p == permission {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, dto.Error(dto.ErrPermissionDenied, "You do not have access to this resource"))
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := GetUserFromContext(c)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error(dto.ErrUnauthorized, "User not authenticated"))
			return
		}

		if user.Role != role {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.Error(dto.ErrRoleNotAllowed, "You do not have access to this resource"))
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Error(dto.ErrUnauthorized, "User not authenticated"))
			return
		}

		if _, ok := allowed[user.Role]; !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.Error(dto.ErrRoleNotAllowed, "You do not have access to this resource"))
			return
		}

		c.Next()
	}
}

func RequireAdmin() gin.HandlerFunc {
	// "admin" and "super_admin" are the roles that can manage RBAC/permissions.
	return RequireAnyRole("admin", "super_admin")
}
