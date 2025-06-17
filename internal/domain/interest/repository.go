package interest

import "context"

type Repository interface {
	GetAll(ctx context.Context, postInterest *[]PostInterest, userID uint, filter GetInterest) (int, error)
	GetDetailByID(ctx context.Context, postInterest *PostInterest, interestID uint) error
	Create(ctx context.Context, interest Interest) (uint, error)
	Delete(ctx context.Context, postID uint, userID uint) (uint, error)
	IsExist(ctx context.Context, userID uint, postID uint) (bool, error)
	IsExistByID(ctx context.Context, interestID uint) (bool, error)
}
