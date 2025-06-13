package warehouse

import "time"

type Warehouse struct {
	ID              uint
	ItemID          uint
	ImportInvoiceID uint
	SenderName      string
	ItemName        string
	SKU             string
	Quantity        int
	Description     string
	Classify        int
	StockPlace      string
	CreatedAt       time.Time
}

type DetailWarehouse struct {
	ID              uint
	ItemID          uint
	ImportInvoiceID uint
	SenderName      string
	ItemName        string
	SKU             string
	Quantity        int
	Description     string
	Classify        int
	StockPlace      string
	ItemWareHouse   []ItemWareHouse
	CreatedAt       time.Time
}

type ItemWareHouse struct {
	ID           uint
	ItemID       uint
	ItemName     string
	CategoryName string
	WarehouseID  uint
	Code         string
	Description  string
	Status       int
}
