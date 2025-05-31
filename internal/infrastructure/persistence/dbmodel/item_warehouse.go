package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type ItemWarehouse struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	ItemID      uint   `gorm:"index"`
	WarehouseID uint   `gorm:"index"`
	Code        string `gorm:"unique;size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// n-1: Mỗi item warehouse thuộc về 1 item và 1 warehouse
	Item      Item      `gorm:"foreignKey:ItemID"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID"`

	// 1-n: Một item warehouse có thể thuộc nhiều post_item_warehouse, appointment_item_warehouse, item_export_invoice
	PostItemWarehouses        []PostItemWarehouse        `gorm:"foreignKey:ItemWarehouseID"`
	AppointmentItemWarehouses []AppointmentItemWarehouse `gorm:"foreignKey:ItemWarehouseID"`
	ItemExportInvoices        []ItemExportInvoice        `gorm:"foreignKey:ItemWarehouseID"`
}
