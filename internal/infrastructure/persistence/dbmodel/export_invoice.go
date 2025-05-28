package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type ExportInvoice struct {
    ID          string         `gorm:"primaryKey;size:255"`
    InvoiceNum  int            `gorm:"unique"`
    AdminID     uint           `gorm:"index"`
    ReceiverID  uint           `gorm:"index"`
    ItemType    string         `gorm:"size:32"`
    ReceiveDate time.Time
    Description string         `gorm:"type:TEXT"`
    IsLock      bool
    IsAnonymous bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
    // Relations
    Admin Admin `gorm:"foreignKey:AdminID"`
    ItemExportInvoices []ItemExportInvoice `gorm:"foreignKey:InvoiceID"`
}