package dbmodel

import (
	importinvoice "final_project/internal/domain/import_invoice"
	"time"

	"gorm.io/gorm"
)

type ImportInvoice struct {
	ID          string `gorm:"primaryKey;size:255"`
	InvoiceNum  int    `gorm:"unique"`
	SenderID    uint   `gorm:"index"`
	ReceiverID  uint   `gorm:"index"`
	Classify    string `gorm:"size:32"` // đổi từ ItemType sang Classify
	Description string `gorm:"type:TEXT"`
	IsLock      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Sender             User                `gorm:"foreignKey:SenderID"`
	Receiver           User                `gorm:"foreignKey:ReceiverID"`
	ItemImportInvoices []ItemImportInvoice `gorm:"foreignKey:InvoiceID;references:ID"`
}

// Domain to DB
func ImportInvoiceDomainToDB(domain importinvoice.ImportInvoice) ImportInvoice {
	var items []ItemImportInvoice

	for _, value := range domain.ItemImportInvoice {
		items = append(items, ItemImportInvoice{
			ItemID:      value.ID,
			Quantity:    value.Quantity,
			Description: value.Description,
		})
	}

	return ImportInvoice{
		InvoiceNum:         domain.InvoiceNum,
		SenderID:           domain.SenderID,
		ReceiverID:         domain.ReceiverID,
		Classify:           domain.Classify,
		Description:        domain.Description,
		IsLock:             domain.IsLock,
		ItemImportInvoices: items,
	}
}
