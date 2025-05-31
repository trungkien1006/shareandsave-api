package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type ItemImportInvoice struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	InvoiceID string `gorm:"size:255;index"`
	ItemID    uint   `gorm:"index"`
	SKU       string `gorm:"unique;size:255"`
	Quantity  int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Invoice ImportInvoice `gorm:"foreignKey:InvoiceID;references:ID"`
	Item    Item          `gorm:"foreignKey:ItemID"`
}
