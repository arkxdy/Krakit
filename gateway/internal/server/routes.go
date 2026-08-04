package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakit/gateway/internal/config"
	"github.com/krakit/gateway/internal/handlers"
	"github.com/krakit/gateway/internal/middleware"
	"github.com/krakit/gateway/internal/proxy/auth"
	"github.com/krakit/gateway/internal/proxy/exam"
)

func RegisterRoutes(router *gin.Engine, cfg config.Config) error {
	router.GET("/health", handlers.Health)
	router.GET("/ready", handlers.Ready)

	// Public auth routes
	apiVersion := router.Group("/api/v1")
	publicAuth := apiVersion.Group("/auth")
	{
		if err := auth.RegisterPublicRoutes(publicAuth, cfg); err != nil {
			return err
		}
	}

	// Protected auth routes
	protectedAuth := apiVersion.Group("/auth")
	protectedAuth.Use(middleware.Auth())
	{
		if err := auth.RegisterProtectedRoutes(protectedAuth, cfg); err != nil {
			return err
		}
	}

	// Exam
	examGroup := apiVersion.Group("/exam")
	examGroup.Use(middleware.Auth())
	{
		if err := exam.RegisterRoutes(examGroup, cfg); err != nil {
			return err
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})

	return nil
}
