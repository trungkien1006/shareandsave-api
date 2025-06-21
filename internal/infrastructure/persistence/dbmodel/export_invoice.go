package dbmodel

import (
	exportinvoice "final_project/internal/domain/export_invoice"
	"time"

	"gorm.io/gorm"
)

type ExportInvoice struct {
	ID          uint   `gorm:"primaryKey;size:255"`
	InvoiceNum  int    `gorm:"unique"`
	SenderID    uint   `gorm:"index"`
	ReceiverID  uint   `gorm:"index"`
	Classify    int    `gorm:"INT"`
	Description string `gorm:"type:TEXT"`
	IsLock      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Sender             User                `gorm:"foreignKey:SenderID"`
	Receiver           User                `gorm:"foreignKey:ReceiverID"`
	ItemExportInvoices []ItemExportInvoice `gorm:"foreignKey:InvoiceID;references:ID"`
}

// Domain to DB
func ExportInvoiceDomainToDB(domain exportinvoice.ExportInvoice) ExportInvoice {
	items := make([]ItemExportInvoice, 0)

	for _, value := range domain.ItemExportInvoices {
		items = append(items, ItemExportInvoice{
			ItemWarehouseID: value.ItemWarehouseID,
			SKU:             value.SKU,
		})
	}

	return ExportInvoice{
		ID:                 domain.ID,
		InvoiceNum:         domain.InvoiceNum,
		SenderID:           domain.SenderID,
		ReceiverID:         domain.ReceiverID,
		Classify:           domain.Classify,
		Description:        domain.Description,
		IsLock:             domain.IsLock,
		ItemExportInvoices: items,
	}
}

// DB to Domain
func ExportInvoiceDBToDomain(db ExportInvoice) exportinvoice.ExportInvoice {
	items := make([]exportinvoice.ItemExportInvoice, 0)

	for _, value := range db.ItemExportInvoices {
		items = append(items, exportinvoice.ItemExportInvoice{
			ID:              value.ID,
			InvoiceID:       value.InvoiceID,
			ItemWarehouseID: value.ItemWarehouseID,
			SKU:             value.SKU,
		})
	}

	return exportinvoice.ExportInvoice{
		ID:           db.ID,
		InvoiceNum:   db.InvoiceNum,
		SenderID:     db.SenderID,
		SenderName:   db.Sender.FullName,
		ReceiverID:   db.ReceiverID,
		ReceiverName: db.Receiver.FullName,
		Classify:     db.Classify,
		Description:  db.Description,
		IsLock:       db.IsLock,
	}
}
