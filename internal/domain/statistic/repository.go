package statistic

import "context"

type Repository interface {
	TotalTransaction(ctx context.Context) (int64, int64, error)
	TotalUser(ctx context.Context, clientID uint) (int64, int64, error)
	TotalPost(ctx context.Context) (int64, int64, error)
	TotalItemWarehouse(ctx context.Context) (int64, int64, error)
}
