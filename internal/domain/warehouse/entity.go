package warehouse

type Warehouse struct {
	ID            uint
	ItemID        uint
	ItemName      string
	SKU           string
	Quantity      int
	Description   string
	Classify      int
	StockPlace    string
	ItemWareHouse []ItemWareHouse
}

type ItemWareHouse struct {
	ID          uint
	ItemID      uint
	ItemName    string
	WarehouseID uint
	Code        string
	Description string
}
