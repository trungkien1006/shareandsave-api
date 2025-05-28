package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	SKU         string `gorm:"unique;size:255"`
	Quantity    int
	Description string `gorm:"type:TEXT"`
	Classify    string `gorm:"size:12"`
	StockPlace  string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	// Relations

	ItemWarehouses []ItemWarehouse `gorm:"foreignKey:WarehouseID"`
}
