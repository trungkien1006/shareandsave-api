package redis

import (
	"context"
	"time"
)

type Repository interface {
	InsertToRedis(ctx context.Context, key string, value string, expiration time.Duration) error
	GetFromRedis(ctx context.Context, key string) (string, error)
	DeleteFromRedis(ctx context.Context, key string) error
	SetToRedisHash(ctx context.Context, hashKey string, field string, value string) error
	GetFromRedisHash(ctx context.Context, hashKey string, field string) (string, error)
	DeleteFromRedisHash(ctx context.Context, hashKey string, fields ...string) error
}
