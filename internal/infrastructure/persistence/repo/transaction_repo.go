package persistence

import (
	"context"
	"final_project/internal/domain/transaction"

	"gorm.io/gorm"
)

type TransactionRepoDB struct {
	db *gorm.DB
}

func NewTransactionRepoDB(db *gorm.DB) *TransactionRepoDB {
	return &TransactionRepoDB{db: db}
}

func (r *TransactionRepoDB) Create(ctx context.Context, transaction *transaction.Transaction) error {
	return nil
}
