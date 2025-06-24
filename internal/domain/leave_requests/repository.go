package leaverequests

import (
	"context"
	"time"
)

type Repository interface {
	IsInLeaveRequest(ctx context.Context, day time.Time) (bool, error)
}
