package redisapp

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

// SetToRedisHash thêm hoặc cập nhật một field vào hash trong Redis
func (r *RedisRepo) SetToRedisHash(ctx context.Context, hashKey string, field string, value string) error {
	err := r.client.HSet(ctx, hashKey, field, value).Err()
	if err != nil {
		return errors.New("Có lỗi khi thêm dữ liệu vào redis hash: " + err.Error())
	}
	return nil
}

func (r *RedisRepo) GetFromRedisHash(ctx context.Context, hashKey string, field string) (string, error) {
	val, err := r.client.HGet(ctx, hashKey, field).Result()
	if err != nil {
		if err == redis.Nil {
			// Field không tồn tại
			return "", nil
		}
		return "", errors.New("Có lỗi khi lấy dữ liệu từ redis hash: " + err.Error())
	}
	return val, nil
}

func (r *RedisRepo) DeleteFromRedisHash(ctx context.Context, hashKey string, fields ...string) error {
	err := r.client.HDel(ctx, hashKey, fields...).Err()
	if err != nil {
		return errors.New("Có lỗi khi xóa field khỏi redis hash: " + err.Error())
	}
	return nil
}
