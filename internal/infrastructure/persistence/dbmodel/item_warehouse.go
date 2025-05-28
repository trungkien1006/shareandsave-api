package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type ItemWarehouse struct {
    ID          uint           `gorm:"primaryKey;autoIncrement"`
    ItemID      uint           `gorm:"index"`
    WarehouseID uint           `gorm:"index"`
    Code        string         `gorm:"unique;size:255"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
    // Relations
    Item      Item      `gorm:"foreignKey:ItemID"`
    Warehouse Warehouse `gorm:"foreignKey:WarehouseID"`
    Requests  []Request `gorm:"foreignKey:ItemWarehouseID"`
    ItemExportInvoices []ItemExportInvoice `gorm:"foreignKey:ItemWarehouseID"`
}