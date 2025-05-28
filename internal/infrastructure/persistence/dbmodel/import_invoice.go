package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type ImportInvoice struct {
    ID          string         `gorm:"primaryKey;size:255"`
    InvoiceNum  int            `gorm:"unique"`
    AdminID     uint           `gorm:"index"`
    SenderID    uint           `gorm:"index"`
    ItemType    string         `gorm:"size:32"`
    SendDate    time.Time
    Description string         `gorm:"type:TEXT"`
    IsLock      bool
    IsAnonymous bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
    // Relations
    Admin Admin `gorm:"foreignKey:AdminID"`
    ItemImportInvoices []ItemImportInvoice `gorm:"foreignKey:InvoiceID"`
}