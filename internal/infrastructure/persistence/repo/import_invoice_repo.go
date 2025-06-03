package persistence

import "gorm.io/gorm"

type ImpoerInvoiceRepoDB struct {
	db *gorm.DB
}

func NewImpoerInvoiceRepoDB(db *gorm.DB) *ImpoerInvoiceRepoDB {
	return &ImpoerInvoiceRepoDB{db: db}
}
