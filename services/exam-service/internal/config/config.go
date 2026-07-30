package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port               string
	DatabaseURL        string
	KafkaBroker        string
	MongoURL           string
	RedisURL           string
	Minio              MinioConfig
	AuthServiceJWKSURL string
}

type MinioConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	UseSSL    bool
}

func LoadConfig() *Config {
	return &Config{
		// Default exam service port (auth uses 8081)
		Port: getEnv("PORT", "8082"),

		// Constructs DB URL: postgres://user:pass@host:port/dbname
		DatabaseURL: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "password"),
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "5433"),
			getEnv("DB_NAME", "exam_db"),
		),

		// For Kafka, remember we use 29092 for local and 9092 for Docker
		KafkaBroker: fmt.Sprintf("%s:%s",
			getEnv("KAFKA_HOST", "localhost"),
			getEnv("KAFKA_PORT", "29092"),
		),

		// For Mongo (root user lives in the admin auth DB)
		MongoURL: fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
			getEnv("MONGO_USER", "admin"),
			getEnv("MONGO_PASS", "password"),
			getEnv("MONGO_HOST", "localhost"),
			getEnv("MONGO_PORT", "27017"),
			getEnv("MONGO_DB", "krakit-question"),
		),

		// For Redis
		RedisURL: redisURL(
			getEnv("REDIS_HOST", "localhost"),
			getEnv("REDIS_PORT", "6379"),
			getEnv("REDIS_PASSWORD", "password"),
		),

		// For Minio
		Minio: MinioConfig{
			Endpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
			AccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
			SecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
			Bucket:    getEnv("MINIO_BUCKET", "exams"),
			UseSSL:    getEnv("MINIO_USE_SSL", "false") == "true",
		},

		AuthServiceJWKSURL: getEnv("AUTH_SERVICE_JWKS", "http://localhost:8081/api/v1/.well-known/jwks.json"),
	}
}

func redisURL(host, port, password string) string {
	if password == "" {
		return fmt.Sprintf("redis://%s:%s", host, port)
	}
	return fmt.Sprintf("redis://:%s@%s:%s", password, host, port)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
