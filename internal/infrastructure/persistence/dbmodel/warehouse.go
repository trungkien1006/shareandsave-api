package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	ItemID      uint   `gorm:"index"` // mới thêm
	SKU         string `gorm:"unique;size:255"`
	Quantity    int
	Description string `gorm:"type:TEXT"`
	Classify    string `gorm:"size:12"`
	StockPlace  string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Item           Item            `gorm:"foreignKey:ItemID"`
	ItemWarehouses []ItemWarehouse `gorm:"foreignKey:WarehouseID"`
}
