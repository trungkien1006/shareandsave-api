package warehouse

import "time"

type Warehouse struct {
	ID              uint
	ItemID          uint
	ImportInvoiceID uint
	SenderName      string
	ReceiverName    string
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
	ReceiverName    string
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
	CreatedAt    time.Time
}

type ItemOldStock struct {
	ItemID            uint
	ItemName          string
	ItemImage         string
	Description       string
	CategoryName      string
	Quantity          uint
	ClaimItemRequests uint
}

type MyClaimRequest struct {
	ItemID       uint
	ItemName     string
	ItemImage    string
	CategoryName string
	Quantity     uint
}

type ClaimRequestItem struct {
	ItemQuantity uint
	Users        []ClaimRequestUser
}

type ClaimRequestUser struct {
	ID       uint
	Quantity uint
}

type CreateClaimRequest struct {
	RequestItems []CreateClaimRequestItem
}

type CreateClaimRequestItem struct {
	ItemID   uint
	Quantity uint
}

type ModifyClaimRequest struct {
	ItemID     uint
	NewQuatity uint
}

type GetItemOldStockFilter struct {
	Page       int
	Limit      int
	Sort       string
	Order      string
	CategoryID int
	Search     string
}

type ItemQuantity struct {
	ItemID   uint  `gorm:"column:item_id"`
	Quantity int64 `gorm:"column:quantity"`
}
