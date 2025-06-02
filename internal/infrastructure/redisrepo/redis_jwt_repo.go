package redisrepo

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	client *redis.Client
}

func NewRedisRepo(c *redis.Client) *RedisRepo {
	return &RedisRepo{client: c}
}

// InsertToRedis insert 1 key-value vào Redis
func (r *RedisRepo) InsertToRedis(ctx context.Context, key string, value string, expiration time.Duration) error {
	if err := r.client.Set(ctx, key, value, expiration).Err(); err != nil {
		errors.New("Có lỗi khi insert vào redis: " + err.Error())
	}
	return nil
}
