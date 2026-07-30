package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krakit/auth-service/internal/cache"
	"github.com/krakit/auth-service/internal/db"
	database "github.com/krakit/auth-service/internal/db/sqlc"
	"github.com/krakit/auth-service/internal/handler"
	"github.com/krakit/auth-service/internal/middleware"
	"github.com/krakit/auth-service/internal/repository"
	"github.com/krakit/auth-service/internal/service"
	"github.com/krakit/auth-service/internal/utils"
)

func main() {
	ctx := context.Background()

	config := utils.LoadConfig()
	keyStore := utils.NewKeyStore()
	jwtMaker := utils.NewJWTMaker(keyStore, utils.JWTIssuer, utils.JWTDuration)

	postgresCfg := db.NewConfig(db.PostgreSQL)
	redisCfg := db.NewConfig(db.Redis)
	// 1. Database Connection
	postgresClient, err := db.NewPostgresClient(postgresCfg)
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL: %v", err)
	}
	defer postgresClient.Close()

	redisClient, err := cache.NewRedisClient(redisCfg)
	if err != nil {
		log.Fatalf("Unable to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// 2. Initialize Layers
	queries := database.New(postgresClient)
	cache := cache.NewRedisCache(redisClient)
	userRepo := repository.NewUserRespository(queries)
	sessionRepo := repository.NewSessionRepository(queries)
	permRepo := repository.NewPerissionRepository(queries)
	authService := service.NewAuthService(userRepo, sessionRepo, permRepo, jwtMaker, config, cache)
	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	permService := service.NewPermissionService(permRepo)
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	sessionHandler := handler.NewSessionHandler(sessionService)
	permissionHandler := handler.NewPermissionHandler(permService)
	jwksHandler := handler.NewJWKSHandler(keyStore)

	//3. Setup Gin
	router := gin.Default()

	//Grouped Routes
	api := router.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})

		authGroup := api.Group("/auth")
		{
			authGroup.POST("/signup", authHandler.Signup)
			authGroup.POST("/login", authHandler.Login)
			authGroup.POST("/google", authHandler.GoogleLogin)
			authGroup.POST("/refresh", authHandler.Refresh)
			authGroup.POST("/logout", authHandler.Logout)
		}

		authAuthedGroup := api.Group("/auth")
		{
			authAuthedGroup.Use(middleware.AuthMiddleware(jwtMaker))
			authAuthedGroup.POST("/logout-all", authHandler.LogoutAll)
		}

		userGroup := api.Group("/users")
		{
			userGroup.Use(middleware.AuthMiddleware(jwtMaker))
			userGroup.GET("/me", userHandler.GetCurrentUser)
			userGroup.PUT("/me", userHandler.UpdateProfile)
			userGroup.POST("/change-password", userHandler.ChangePassword)
		}

		sessions := api.Group("/sessions")
		{
			sessions.Use(middleware.AuthMiddleware(jwtMaker))
			sessions.GET("", sessionHandler.ListSessions)
			sessions.DELETE("/:id", sessionHandler.RevokeSession)
		}

		permissions := api.Group("/permissions")
		{
			permissions.Use(middleware.AuthMiddleware(jwtMaker))
			permissions.Use(middleware.RequireAdmin())
			permissions.GET("", permissionHandler.ListPermissions)
		}

		roles := api.Group("/roles")
		{
			roles.Use(middleware.AuthMiddleware(jwtMaker))
			roles.Use(middleware.RequireAdmin())
			roles.POST("/:role/permissions", permissionHandler.AssignPermissionToRole)
			roles.DELETE("/:role/permissions/:permission_id", permissionHandler.RemovePermissionFromRole)
		}

		api.GET("/.well-known/jwks.json", jwksHandler.GetJWKS)
	}

	//4. Graceful Start
	port := os.Getenv("AUTH_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// 5. Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Printf("Auth Service (Gin) running on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen Error: %s\n", err)
		}
	}()

	// 6. Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so no need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
