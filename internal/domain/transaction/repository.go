package transaction

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, transaction *Transaction) error
	// CheckPostItemQuantityOver(ctx context.Context, postID uint, quantity int) error
}
