package redis

import (
	"context"
	"time"
)

type Reposity interface {
	InsertToRedis(ctx context.Context, key string, value string, expiration time.Duration) error
	GetFromRedis(ctx context.Context, key string) (string, error)
	DeleteFromRedis(ctx context.Context, key string) error
}
