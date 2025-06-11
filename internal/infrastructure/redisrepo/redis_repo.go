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

// GetFromRedis lấy giá trị từ Redis theo key
func (r *RedisRepo) GetFromRedis(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// key không tồn tại trong Redis
			return "", nil
		}
		return "", errors.New("Có lỗi khi lấy dữ liệu từ redis: " + err.Error())
	}
	return val, nil
}

// DeleteFromRedis xóa một key khỏi Redis
func (r *RedisRepo) DeleteFromRedis(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return errors.New("Có lỗi khi xóa key khỏi redis: " + err.Error())
	}
	return nil
}
