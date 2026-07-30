package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/krakit/exam-service/internal/config"
	"github.com/krakit/exam-service/internal/db"
	database "github.com/krakit/exam-service/internal/db/sqlc"
	"github.com/krakit/exam-service/internal/handler"
	"github.com/krakit/exam-service/internal/middleware"
	"github.com/krakit/exam-service/internal/repository"
	"github.com/krakit/exam-service/internal/service"
	"github.com/krakit/exam-service/internal/worker"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Init Connections
	conns, err := db.NewConnections(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize connections: %v", err)
	}
	defer conns.Close()

	pool := worker.NewPool(5) // 5 concurrent workers
	defer pool.Shutdown()     // waits for in-flight jobs before exit

	jwksClient := middleware.NewJWKSClient(cfg.AuthServiceJWKSURL)
	queries := database.New(conns.Postgres.Pool)
	ansRepo := repository.NewAnswerRepository(queries)
	attemptQuesMapRepo := repository.NewAttemptQuestionMapRepository(queries)
	attemptRepo := repository.NewAttemptRepository(queries)
	examRepo := repository.NewExamRepository(queries)
	quesRepo := repository.NewQuestionRepository(conns.Mongo)
	quesSetRepo := repository.NewQuestionSetsRepository(queries)
	sectionRepo := repository.NewSectionRepository(queries)
	sectionScoreRepo := repository.NewSectionScoresRepository(queries)
	settingRepo := repository.NewExamSettingsRepository(queries)
	subjectRepo := repository.NewSubjectRepository(queries)

	examSvc := service.NewExamService(&examRepo, &settingRepo, &sectionRepo, &subjectRepo)
	attemptSvc := service.NewAttemptService(&attemptRepo, &attemptQuesMapRepo, &ansRepo, &quesRepo, &quesSetRepo, &sectionScoreRepo, pool)
	questionSvc := service.NewQuestionService(quesRepo, quesSetRepo)

	examHandler := handler.NewExamHandler(examSvc)
	attemptHandler := handler.NewAttemptHandler(attemptSvc)
	questionHandler := handler.NewQuestionHandler(questionSvc)

	//3. Setup Gin
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})

		// admin routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(jwksClient))
		admin.Use(middleware.RequireAdmin())
		{
			// exams
			admin.POST("/exams", examHandler.CreateExam)
			admin.PUT("/exams/:id", examHandler.UpdateExam)
			admin.DELETE("/exams/:id", examHandler.DisableExam)
			admin.POST("/exams/:id/publish", examHandler.PublishExam)
			admin.POST("/exams/:id/settings", examHandler.CreateExamSettings)

			// sections
			admin.POST("/exams/:id/sections", examHandler.CreateSection)
			admin.PUT("/exams/:id/sections/:section_id", examHandler.UpdateSection)

			// subjects
			admin.POST("/subjects", examHandler.CreateSubject)

			// question sets
			admin.POST("/exams/:id/question-sets", questionHandler.AttachQuestionSet)

			// questions
			admin.POST("/questions", questionHandler.CreateQuestions)
			admin.PUT("/questions/bulk", questionHandler.BulkUpsertQuestions)
			admin.DELETE("/questions/set/:set_id", questionHandler.DeleteQuestionsBySet)

			// passages
			admin.POST("/passages", questionHandler.CreatePassages)
			admin.PUT("/passages/bulk", questionHandler.BulkUpsertPassages)
			admin.DELETE("/passages/section/:section_id", questionHandler.DeletePassagesBySection)
		}

		// user routes
		exams := api.Group("/exams")
		exams.Use(middleware.AuthMiddleware(jwksClient))
		{
			exams.GET("", examHandler.ListExams)
			exams.GET("/paginated", examHandler.ListExamsPaginated)
			exams.GET("/:id/sections", examHandler.GetSections)
			exams.GET("/question-sets/:id", questionHandler.GetQuestionSets)
		}

		subjects := api.Group("/subjects")
		subjects.Use(middleware.AuthMiddleware(jwksClient))
		{
			subjects.GET("", examHandler.GetSubjects)
		}

		attempts := api.Group("/attempts")
		attempts.Use(middleware.AuthMiddleware(jwksClient))
		{
			attempts.POST("", attemptHandler.StartAttempt)
			attempts.GET("/my", attemptHandler.GetUserAttempts)
			attempts.GET("/:id", attemptHandler.GetAttempt)
			attempts.POST("/:id/submit", attemptHandler.SubmitAttempt)
			attempts.GET("/:id/result", attemptHandler.GetResult)
			attempts.GET("/:id/questions", attemptHandler.GetAttemptQuestions)
			attempts.GET("/:id/sections/:section_id/questions", attemptHandler.GetAttemptQuestionsBySection)
			attempts.POST("/:id/answers", attemptHandler.SaveAnswer)
			attempts.GET("/:id/answers", attemptHandler.GetAnswers)
		}
	}

	// 4. Start HTTP server
	go func() {
		log.Printf("Exam Service (Gin) running on port %s", cfg.Port)
		if err := router.Run(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen Error: %s\n", err)
		}
	}()

	// Wait for termination signal (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down Exam Service...")
}
