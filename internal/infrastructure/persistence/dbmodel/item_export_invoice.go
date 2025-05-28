package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type ItemExportInvoice struct {
    ID              uint           `gorm:"primaryKey;autoIncrement"`
    InvoiceID       string         `gorm:"index"`
    ItemWarehouseID uint           `gorm:"index"`
    SKU             string         `gorm:"unique;size:255"`
    Quantity        int8
    CreatedAt       time.Time
    UpdatedAt       time.Time
    DeletedAt       gorm.DeletedAt `gorm:"index"`
    // Relations
    Invoice        ExportInvoice  `gorm:"foreignKey:InvoiceID"`
    ItemWarehouse  ItemWarehouse  `gorm:"foreignKey:ItemWarehouseID"`
}