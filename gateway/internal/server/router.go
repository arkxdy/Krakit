package server

import (
	"github.com/gin-gonic/gin"
	"github.com/krakit/gateway/internal/config"
	"github.com/krakit/gateway/internal/middleware"
)

func NewRouter(cfg config.Config) *gin.Engine {
	router := gin.New()

	router.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
	)

	RegisterRoutes(router, cfg)

	return router
}
