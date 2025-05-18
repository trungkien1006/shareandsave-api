package warehouse

import (
	"final-project/internal/domain/item"
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	ItemID      uint   `gorm:"index"`
	SKU         string `gorm:"size:255;unique"`
	Quantity    int
	Description string `gorm:"type:TEXT"`
	Classify    string `gorm:"size:12"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Item item.Item `gorm:"foreignKey:ItemID"`
}

func NewWarehouse(itemID uint, sku string, quantity int, description, classify string) *Warehouse {
	return &Warehouse{
		ItemID:      itemID,
		SKU:         sku,
		Quantity:    quantity,
		Description: description,
		Classify:    classify,
	}
}
