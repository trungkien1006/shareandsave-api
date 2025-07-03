package statistic

import "context"

type Repository interface {
	TotalTransaction(ctx context.Context) (int64, int64, error)
}
