package dbmodel

import (
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/domain/warehouse"
	"time"

	"gorm.io/gorm"
)

type ImportInvoice struct {
	ID          uint   `gorm:"primaryKey;size:255"`
	InvoiceNum  int    `gorm:"unique"`
	SenderID    uint   `gorm:"index"`
	ReceiverID  uint   `gorm:"index"`
	Classify    int    `gorm:"type:int"` // đổi từ ItemType sang Classify
	Description string `gorm:"type:TEXT"`
	IsLock      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Sender             User                `gorm:"foreignKey:SenderID"`
	Receiver           User                `gorm:"foreignKey:ReceiverID"`
	ItemImportInvoices []ItemImportInvoice `gorm:"foreignKey:InvoiceID;references:ID"`
	Warehouses         []Warehouse         `gorm:"foreignKey:ImportInvoiceID;references:ID"`
}

// Domain to DB
func ImportInvoiceDomainToDB(domain importinvoice.ImportInvoice) ImportInvoice {
	var (
		warehouses []Warehouse
		items      []ItemImportInvoice
	)

	for _, value := range domain.ItemImportInvoice {
		items = append(items, ItemImportInvoice{
			ItemID:      value.ItemID,
			Quantity:    value.Quantity,
			Description: value.Description,
		})
	}

	for _, v := range domain.Warehouses {
		var itemWarehouses []ItemWarehouse

		for _, value := range v.ItemWareHouse {
			itemWarehouses = append(itemWarehouses, ItemWarehouse{
				ItemID:      value.ItemID,
				Code:        value.Code,
				Description: value.Description,
				Status:      value.Status,
			})
		}

		warehouses = append(warehouses, Warehouse{
			ItemID:         v.ItemID,
			SKU:            v.SKU,
			Quantity:       v.Quantity,
			Classify:       v.Classify,
			Description:    v.Description,
			StockPlace:     v.StockPlace,
			ItemWarehouses: itemWarehouses,
		})
	}

	return ImportInvoice{
		InvoiceNum:         domain.InvoiceNum,
		SenderID:           domain.SenderID,
		ReceiverID:         domain.ReceiverID,
		Classify:           domain.Classify,
		Description:        domain.Description,
		IsLock:             domain.IsLock,
		ItemImportInvoices: items,
		Warehouses:         warehouses,
	}
}

// Domain to DB
func ImportInvoiceDBToDomain(db ImportInvoice) importinvoice.ImportInvoice {
	var (
		warehouses []warehouse.DetailWarehouse
		items      []importinvoice.ItemImportInvoice
	)

	for _, value := range db.ItemImportInvoices {
		items = append(items, importinvoice.ItemImportInvoice{
			ItemID:      value.ItemID,
			Quantity:    value.Quantity,
			Description: value.Description,
		})
	}

	for _, v := range db.Warehouses {
		var itemWarehouses []warehouse.ItemWareHouse

		for _, value := range v.ItemWarehouses {
			itemWarehouses = append(itemWarehouses, warehouse.ItemWareHouse{
				ItemID:      value.ItemID,
				ItemName:    value.Item.Name,
				Code:        value.Code,
				Description: value.Description,
				Status:      value.Status,
			})
		}

		warehouses = append(warehouses, warehouse.DetailWarehouse{
			ItemID:        v.ItemID,
			SKU:           v.SKU,
			Quantity:      v.Quantity,
			Classify:      v.Classify,
			Description:   v.Description,
			StockPlace:    v.StockPlace,
			ItemWareHouse: itemWarehouses,
		})
	}

	return importinvoice.ImportInvoice{
		InvoiceNum:        db.InvoiceNum,
		SenderID:          db.SenderID,
		ReceiverID:        db.ReceiverID,
		Classify:          db.Classify,
		Description:       db.Description,
		IsLock:            db.IsLock,
		ItemImportInvoice: items,
		Warehouses:        warehouses,
	}
}
