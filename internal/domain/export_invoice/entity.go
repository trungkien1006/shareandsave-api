package exportinvoice

import (
	"final_project/internal/domain/admin"
	"os/user"
	"time"

	"gorm.io/gorm"
)

type ExportInvoice struct {
	ID          string `gorm:"primaryKey;size:255"`
	InvoiceNum  int
	AdminID     uint   `gorm:"index"`
	ReceiverID  uint   `gorm:"index"`
	ItemType    string `gorm:"size:32"`
	ReceiveDate time.Time
	Description string `gorm:"type:TEXT"`
	IsLock      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Admin       admin.Admin    `gorm:"foreignKey:AdminID"`
	Receiver    user.User      `gorm:"foreignKey:ReceiverID"`
}

func NewExportInvoice(id string, invoiceNum int, adminID, receiverID uint, itemType string, receiveDate time.Time, description string, isLock bool) *ExportInvoice {
	return &ExportInvoice{
		ID:          id,
		InvoiceNum:  invoiceNum,
		AdminID:     adminID,
		ReceiverID:  receiverID,
		ItemType:    itemType,
		ReceiveDate: receiveDate,
		Description: description,
		IsLock:      isLock,
	}
}
