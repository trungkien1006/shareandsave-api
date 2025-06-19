package warehousedto

import "time"

type WarehouseDTO struct {
	ID              uint      `json:"id"`
	ItemID          uint      `json:"itemID"`
	ImportInvoiceID uint      `json:"importInvoiceID"`
	SenderName      string    `json:"senderName"`
	ReceiverName    string    `json:"receiverName"`
	ItemName        string    `json:"itemName"`
	SKU             string    `json:"SKU"`
	Quantity        int       `json:"quantity"`
	Description     string    `json:"description"`
	Classify        int       `json:"classify"`
	StockPlace      string    `json:"stockPlace"`
	CreatedAt       time.Time `json:"createdAt"`
}

type DetailWarehouseDTO struct {
	ID              uint               `json:"id"`
	ItemID          uint               `json:"itemID"`
	ImportInvoiceID uint               `json:"importInvoiceID"`
	SenderName      string             `json:"senderName"`
	ReceiverName    string             `json:"receiverName"`
	ItemName        string             `json:"itemName"`
	SKU             string             `json:"SKU"`
	Quantity        int                `json:"quantity"`
	Description     string             `json:"description"`
	Classify        int                `json:"classify"`
	StockPlace      string             `json:"stockPlace"`
	ItemWareHouse   []ItemWareHouseDTO `json:"itemWarehouses"`
	CreatedAt       time.Time          `json:"createdAt"`
}

type ItemWareHouseDTO struct {
	ID           uint   `json:"id"`
	ItemID       uint   `json:"itemID"`
	ItemName     string `json:"itemName"`
	CategoryName string `json:"categoryName"`
	WarehouseID  uint   `json:"warehouseID"`
	Code         string `json:"code"`
	Description  string `json:"description"`
	Status       int    `json:"status"`
}

type ItemOldStockDTO struct {
	ItemID            uint   `json:"item_id"`
	ItemName          string `json:"item_name"`
	ItemImage         string `json:"item_image"`
	Description       string `json:"description"`
	CategoryName      string `json:"category_name"`
	Quantity          uint   `json:"quantity"`
	ClaimItemRequests uint   `json:"claim_item_requests"`
}
