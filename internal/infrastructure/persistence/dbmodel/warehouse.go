package dbmodel

import (
	"final_project/internal/domain/warehouse"
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	ItemID      uint   `gorm:"index"` // mới thêm
	SKU         string `gorm:"unique;size:255"`
	Quantity    int
	Description string `gorm:"type:TEXT"`
	Classify    int    `gorm:"type:int"`
	StockPlace  string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Item           Item            `gorm:"foreignKey:ItemID"`
	ItemWarehouses []ItemWarehouse `gorm:"foreignKey:WarehouseID"`
}

// Domain to DB
func WarehouseDomainToDB(domain warehouse.Warehouse) Warehouse {
	var items []ItemWarehouse

	for _, value := range domain.ItemWareHouse {
		items = append(items, ItemWarehouse{
			ItemID:      value.ItemID,
			Code:        value.Code,
			Description: value.Description,
		})
	}

	return Warehouse{
		ItemID:         domain.ItemID,
		SKU:            domain.SKU,
		Quantity:       domain.Quantity,
		Classify:       domain.Classify,
		Description:    domain.Description,
		StockPlace:     domain.StockPlace,
		ItemWarehouses: items,
	}
}

// Domain to DB
func WarehouseDBToDomain(db Warehouse, itemName string) warehouse.Warehouse {
	var items []warehouse.ItemWareHouse

	for _, value := range db.ItemWarehouses {
		items = append(items, warehouse.ItemWareHouse{
			ID:          value.ID,
			ItemID:      value.ItemID,
			ItemName:    itemName,
			WarehouseID: value.WarehouseID,
			Code:        value.Code,
			Description: value.Description,
		})
	}

	return warehouse.Warehouse{
		ID:            db.ID,
		ItemID:        db.ItemID,
		ItemName:      itemName,
		SKU:           db.SKU,
		Quantity:      db.Quantity,
		Classify:      db.Classify,
		Description:   db.Description,
		StockPlace:    db.StockPlace,
		ItemWareHouse: items,
	}
}
