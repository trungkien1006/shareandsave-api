package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type ItemExportInvoice struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	InvoiceID       int    `gorm:"index"`
	ItemWarehouseID uint   `gorm:"index"`
	SKU             string `gorm:"unique;size:255"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	Invoice       ExportInvoice `gorm:"foreignKey:InvoiceID;references:ID"`
	ItemWarehouse ItemWarehouse `gorm:"foreignKey:ItemWarehouseID"`
}
