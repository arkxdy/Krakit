package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/krakit/gateway/internal/config"
	proxy "github.com/krakit/gateway/internal/proxy"
)

func RegisterPublicRoutes(r *gin.RouterGroup, cfg config.Config) error {
	p, err := proxy.New(cfg.Services.Auth)
	if err != nil {
		return err
	}

	r.POST("/login", gin.WrapH(p))
	r.POST("/register", gin.WrapH(p))
	r.POST("/refresh", gin.WrapH(p))
	r.POST("/forgot-password", gin.WrapH(p))
	r.POST("/reset-password", gin.WrapH(p))

	return nil
}

func RegisterProtectedRoutes(r *gin.RouterGroup, cfg config.Config) error {
	p, err := proxy.New(cfg.Services.Auth)
	if err != nil {
		return err
	}

	r.POST("/logout", gin.WrapH(p))
	r.GET("/me", gin.WrapH(p))
	r.PUT("/password", gin.WrapH(p))
	r.POST("/verify-email", gin.WrapH(p))

	return nil
}
