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
		// Default to 8081 for Auth, 8082 for Exam
		Port: getEnv("PORT", "8081"),

		// Constructs DB URL: postgres://user:pass@host:port/dbname
		DatabaseURL: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			getEnv("DB_USER", "admin"),
			getEnv("DB_PASSWORD", "password"),
			getEnv("DB_HOST", "localhost"), // Use 'auth-db' when in Docker
			getEnv("DB_PORT", "5432"),
			getEnv("DB_NAME", "krakit_auth"),
		),

		// For Kafka, remember we use 29092 for local and 9092 for Docker
		KafkaBroker: fmt.Sprintf("%s:%s",
			getEnv("KAFKA_HOST", "localhost"),
			getEnv("KAFKA_PORT", "29092"),
		),

		// For Mongo
		MongoURL: fmt.Sprintf("mongodb://%s:%s@%s:%s",
			getEnv("MONGO_USER", "admin"),
			getEnv("MONGO_PASS", "password"),
			getEnv("MONGO_HOST", "localhost"),
			getEnv("MONGO_PORT", "27017"),
		),

		// For Redis
		RedisURL: fmt.Sprintf("redis://%s:%s",
			getEnv("REDIS_HOST", "localhost"),
			getEnv("REDIS_PORT", "6379"),
		),

		// For Minio
		Minio: MinioConfig{
			Endpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
			AccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
			SecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
			Bucket:    getEnv("MINIO_BUCKET", "exams"),
			UseSSL:    getEnv("MINIO_USE_SSL", "false") == "true",
		},

		AuthServiceJWKSURL: getEnv("AUTH_SERVICE_JWKS", "hhttp://localhost:8081/api/v1/.well-known/jwks.json"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
