package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	KafkaBroker string
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
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
