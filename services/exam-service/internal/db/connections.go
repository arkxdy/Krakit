package db

import (
	"fmt"

	"github.com/krakit/exam-service/internal/config"
)

type Connections struct {
	Postgres *PostgresClient
	Mongo    *MongoClient
	Bucket   *BucketClient
	Redis    *RedisClient
	Kafka    *KafkaClient
}

func NewConnections(cfg *config.Config) (*Connections, error) {
	pg, err := connectPostgres(cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	mongo, err := connectMongo(cfg.MongoURL)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	minio, err := connectMinio(cfg.Minio)
	if err != nil {
		return nil, fmt.Errorf("minio: %w", err)
	}

	rdb, err := connectRedis(cfg.RedisURL)
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}

	kfk, err := connectKafka(cfg.KafkaBroker)
	if err != nil {
		return nil, fmt.Errorf("kafka: %w", err)
	}

	return &Connections{pg, mongo, minio, rdb, kfk}, nil
}

func (c *Connections) Close() {
	c.Postgres.Close()
	c.Mongo.Close()
	// minio client doesn't need explicit close
}
