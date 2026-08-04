package exam

import (
	"github.com/gin-gonic/gin"

	"github.com/krakit/gateway/internal/config"
	proxyutil "github.com/krakit/gateway/internal/proxy"
)

func RegisterRoutes(r *gin.RouterGroup, cfg config.Config) error {
	p, err := proxyutil.New(cfg.Services.Auth)
	if err != nil {
		return err
	}

	exam := r.Group("/exam")
	{
		exam.GET("/questions", gin.WrapH(p))
		exam.POST("/submit", gin.WrapH(p))
	}

	return nil
}
