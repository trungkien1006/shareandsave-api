package reference

import (
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/domain/item"
	"time"

	"gorm.io/gorm"
)

type ItemImportInvoice struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	InvoiceID  string `gorm:"size:255;index"`
	ItemID     uint   `gorm:"index"`
	SKU        string `gorm:"size:255;unique"`
	Quantity   uint8
	StockPlace *string `gorm:"size:255;default:null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt              `gorm:"index"`
	Invoice    importinvoice.ImportInvoice `gorm:"foreignKey:InvoiceID"`
	Item       item.Item                   `gorm:"foreignKey:ItemID"`
}
