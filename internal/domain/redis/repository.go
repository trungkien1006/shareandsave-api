package redis

import (
	"context"
	"time"
)

type Reposity interface {
	InsertToRedis(ctx context.Context, key string, value string, expiration time.Duration) error
}
