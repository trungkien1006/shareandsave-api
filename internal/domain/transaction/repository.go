package transaction

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context, transactions *[]DetailTransaction, req FilterTransaction) (int, error)
	GetByID(ctx context.Context, transactionID uint, transaction *Transaction) error
	GetDetailPendingTransaction(ctx context.Context, transaction *DetailTransaction, interestID uint) error
	Create(ctx context.Context, transaction *Transaction) error
	Update(ctx context.Context, transaction *Transaction) error
	// CheckPostItemQuantityOver(ctx context.Context, postID uint, quantity int) error
	IsExist(ctx context.Context, transactionID uint) (bool, error)
}
