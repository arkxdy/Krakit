package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/krakit/exam-service/internal/config"
	"github.com/krakit/exam-service/internal/db"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Init DB
	pgClient, err := db.NewPostgresClient(cfg)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pgClient.Close()

	// 3. Keep service running
	log.Printf("Auth Service started on port %s", cfg.Port)

	// Wait for termination signal (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down Exam Service...")
}
