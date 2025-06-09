package transaction

import (
	"context"
)

type Repository interface {
	GetByID(ctx context.Context, transactionID uint, transaction *Transaction) error
	Create(ctx context.Context, transaction *Transaction) error
	Update(ctx context.Context, transaction *Transaction) error
	// CheckPostItemQuantityOver(ctx context.Context, postID uint, quantity int) error
	IsExist(ctx context.Context, transactionID uint) (bool, error)
}
