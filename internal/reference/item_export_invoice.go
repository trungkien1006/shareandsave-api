package reference

import (
	exportinvoice "final_project/internal/domain/export_invoice"
	"final_project/internal/domain/item"
	"time"

	"gorm.io/gorm"
)

type ItemExportInvoice struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	InvoiceID string `gorm:"size:255;index"`
	ItemID    uint   `gorm:"index"`
	SKU       string `gorm:"size:255;unique"`
	Quantity  uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt              `gorm:"index"`
	Invoice   exportinvoice.ExportInvoice `gorm:"foreignKey:InvoiceID"`
	Item      item.Item                   `gorm:"foreignKey:ItemID"`
}
