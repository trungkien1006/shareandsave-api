package interest

import "context"

type Repository interface {
	GetAll(ctx context.Context, postInterest *[]PostInterest, userID uint, filter GetInterest) (int, error)
}
