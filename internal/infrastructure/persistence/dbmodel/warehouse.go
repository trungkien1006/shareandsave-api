package dbmodel

import (
	"final_project/internal/domain/warehouse"
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	ImportInvoiceID uint   `gorm:"index"`
	ItemID          uint   `gorm:"index"` // mới thêm
	SKU             string `gorm:"unique;size:255"`
	Quantity        int
	Description     string `gorm:"type:TEXT"`
	Classify        int    `gorm:"type:int"`
	StockPlace      string `gorm:"size:255"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	ImportInvoice  ImportInvoice   `gorm:"foreignKey:ImportInvoiceID"`
	Item           Item            `gorm:"foreignKey:ItemID"`
	ItemWarehouses []ItemWarehouse `gorm:"foreignKey:WarehouseID"`
}

type DetailWarehouse struct {
	Warehouse
	ItemName     string `gorm:"column:item_name"`
	SenderName   string `gorm:"column:sender_name"`
	ReceiverName string `gorm:"column:receiver_name"`
}

// DB to Domain
func DetailDBToDomain(db DetailWarehouse) warehouse.Warehouse {
	return warehouse.Warehouse{
		ID:              db.ID,
		ItemID:          db.ItemID,
		ImportInvoiceID: db.ImportInvoiceID,
		SenderName:      db.SenderName,
		ReceiverName:    db.ReceiverName,
		ItemName:        db.ItemName,
		SKU:             db.SKU,
		Quantity:        db.Quantity,
		Description:     db.Description,
		Classify:        db.Classify,
		StockPlace:      db.StockPlace,
		CreatedAt:       db.CreatedAt,
	}
}

// DB to Domain
func ItemWarehouseDBToDomain(db ItemWarehouse) warehouse.ItemWareHouse {
	return warehouse.ItemWareHouse{
		ID:           db.ID,
		ItemID:       db.ItemID,
		ItemName:     db.Item.Name,
		CategoryName: db.Item.Category.Name,
		WarehouseID:  db.WarehouseID,
		Code:         db.Code,
		Description:  db.Description,
		Status:       db.Status,
	}
}

// Domain to DB
func UpdateDomainToDB(domain warehouse.DetailWarehouse) Warehouse {
	var itemWarehouses []ItemWarehouse

	for _, value := range domain.ItemWareHouse {
		itemWarehouses = append(itemWarehouses, ItemWarehouse{
			ID:          value.ID,
			Description: value.Description,
		})
	}

	return Warehouse{
		ID:             domain.ID,
		Description:    domain.Description,
		StockPlace:     domain.StockPlace,
		ItemWarehouses: itemWarehouses,
	}
}

// DB to Domain
func DetailDBToDetailDomain(db DetailWarehouse) warehouse.DetailWarehouse {
	itemWarehouses := make([]warehouse.ItemWareHouse, 0)

	for _, value := range db.ItemWarehouses {
		itemWarehouses = append(itemWarehouses, warehouse.ItemWareHouse{
			ID:           value.ID,
			ItemID:       value.ItemID,
			ItemName:     value.Item.Name,
			CategoryName: value.Item.Category.Name,
			WarehouseID:  value.WarehouseID,
			Code:         value.Code,
			Description:  value.Description,
			Status:       value.Status,
		})
	}

	return warehouse.DetailWarehouse{
		ID:              db.ID,
		ItemID:          db.ItemID,
		ImportInvoiceID: db.ImportInvoiceID,
		SenderName:      db.SenderName,
		ReceiverName:    db.ReceiverName,
		ItemName:        db.ItemName,
		SKU:             db.SKU,
		Quantity:        db.Quantity,
		Description:     db.Description,
		Classify:        db.Classify,
		StockPlace:      db.StockPlace,
		ItemWareHouse:   itemWarehouses,
		CreatedAt:       db.CreatedAt,
	}
}

// Domain to DB
func WarehouseDomainToDB(domain warehouse.DetailWarehouse) Warehouse {
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
func WarehouseDBToDomain(db Warehouse, itemName string) warehouse.DetailWarehouse {
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

	return warehouse.DetailWarehouse{
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
