package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type ExportInvoice struct {
	ID          string `gorm:"primaryKey;size:255"`
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
