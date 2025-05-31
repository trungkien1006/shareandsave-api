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
	ItemType    string `gorm:"size:32"`
	ReceiveDate time.Time
	Description string `gorm:"type:TEXT"`
	IsLock      bool
	IsAnonymous bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Sender             User                `gorm:"foreignKey:SenderID"`
	Receiver           User                `gorm:"foreignKey:ReceiverID"`
	ItemExportInvoices []ItemExportInvoice `gorm:"foreignKey:InvoiceID;references:ID"`
}
