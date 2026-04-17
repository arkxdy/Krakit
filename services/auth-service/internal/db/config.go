package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBType string

const (
	PostgreSQL DBType = "postgres"
	Redis      DBType = "redis"
)

type Database struct {
	Type     DBType
	Postgres *pgxpool.Pool
	Redis    *redis.Client
}

type Config struct {
	Type     DBType
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewConfig(dbType DBType) *Config {
	switch dbType {
	case PostgreSQL:
		return &Config{
			Type:     dbType,
			Host:     getEnv("DB_HOST", "db"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "auth_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		}
	case Redis:
		return &Config{
			Type:     dbType,
			Host:     getEnv("REDIS_HOST", "redis"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DBName:   getEnv("REDIS_DB", "0"),
		}
	default:
		log.Fatalf("unsupported database type: %s", dbType)
		return nil
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (db *Database) Close() error {
	switch db.Type {
	case PostgreSQL:
		db.Postgres.Close()
		return nil
	case Redis:
		return db.Redis.Close()
	default:
		return fmt.Errorf("unsupported database type: %s", db.Type)
	}
}
