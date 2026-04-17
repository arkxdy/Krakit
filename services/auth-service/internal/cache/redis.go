package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/krakit/auth-service/internal/db"
)

func NewRedisClient(cfg *db.Config) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
	}), nil
}
