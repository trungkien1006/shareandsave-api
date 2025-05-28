package importinvoice

import (
	"time"

	"gorm.io/gorm"
)

type ImportInvoice struct {
	ID          string `gorm:"primaryKey;size:255"`
	InvoiceNum  int
	AdminID     uint   `gorm:"index"`
	SenderID    uint   `gorm:"index"`
	ItemType    string `gorm:"size:32"`
	SendDate    time.Time
	Description string `gorm:"type:TEXT"`
	IsLock      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func NewImportInvoice(id string, invoiceNum int, adminID, senderID uint, itemType string, sendDate time.Time, description string, isLock bool) *ImportInvoice {
	return &ImportInvoice{
		ID:          id,
		InvoiceNum:  invoiceNum,
		AdminID:     adminID,
		SenderID:    senderID,
		ItemType:    itemType,
		SendDate:    sendDate,
		Description: description,
		IsLock:      isLock,
	}
}
