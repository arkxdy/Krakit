package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Exists(ctx context.Context, key string) (bool, error)
	Delete(ctx context.Context, key string) error
	Increment(ctx context.Context, key string, ttl time.Duration) (int64, error)
	SetNX(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error)
}

type redisCache struct {
	client *redis.Client
}

// Increment implements [Cache].
func (r *redisCache) Increment(ctx context.Context, key string, ttl time.Duration) (int64, error) {
	val, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if val == 1 {
		r.client.Expire(ctx, key, ttl)
	}

	return val, nil
}

// SetNX implements [Cache].
func (r *redisCache) SetNX(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error) {
	return r.client.SetNX(ctx, key, value, ttl).Result()
}

// Delete implements [Cache].
func (r *redisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Exists implements [Cache].
func (r *redisCache) Exists(ctx context.Context, key string) (bool, error) {
	res, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

// Get implements [Cache].
func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // cache miss
	}
	return val, err
}

// Set implements [Cache].
func (r *redisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func NewRedisCache(client *redis.Client) Cache {
	return &redisCache{client: client}
}
