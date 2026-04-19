package db

import (
	"github.com/go-redis/redis"
)

type RedisClient struct {
	redis *redis.Client
}

func connectRedis(url string) (*RedisClient, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opts)
	if err := rdb.Ping().Err(); err != nil {
		return nil, err
	}
	return &RedisClient{redis: rdb}, nil
}

func (c *RedisClient) Close() {
	c.redis.Close()
}
