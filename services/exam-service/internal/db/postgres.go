package db

import (
	"context"
	"log"
	"time"

	"github.com/krakit/exam-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClient struct {
	Pool *pgxpool.Pool
}

func NewPostgresClient(cfg *config.Config) (*PostgresClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Parse the connection string from our config utility
	dbConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	// Senior Tip: Configure the pool behavior for production
	dbConfig.MaxConns = 25
	dbConfig.MinConns = 5
	dbConfig.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to Exam Postgres")
	return &PostgresClient{Pool: pool}, nil
}

func (c *PostgresClient) Close() {
	c.Pool.Close()
}
