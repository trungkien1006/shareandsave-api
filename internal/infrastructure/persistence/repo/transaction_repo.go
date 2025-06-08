package persistence

import "gorm.io/gorm"

type TransactionRepoDB struct {
	db *gorm.DB
}

func NewTransactionRepoDB(db *gorm.DB) *TransactionRepoDB {
	return &TransactionRepoDB{db: db}
}
