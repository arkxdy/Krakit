package config

import "os"

type Services struct {
	Auth string
	Exam string
}

func LoadServices() Services {
	return Services{
		Auth: getEnv("AUTH_SERVICE_URL", "http://localhost:8081"),
		Exam: getEnv("EXAM_SERVICE_URL", "http://localhost:8082"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
