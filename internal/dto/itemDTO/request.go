package itemDTO

type GetAllItemRequest struct {
	Page   int    `query:"page" binding:"omitempty,min=1"`
	Limit  int    `query:"limit" binding:"omitempty,min=8"`
	Sort   string `query:"sort" binding:"omitempty,oneof=ASC DESC"`
	Order  string `query:"order"`
	Filter string `query:"filter"`
}

// Request dùng cho lấy 1 item theo ID
type GetItemByIDRequest struct {
	ItemID uint `uri:"itemID" binding:"required"`
}

type CreateItemRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type UpdateItemRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// Request dùng để delete item theo ID
type DeleteItemRequest struct {
	ItemID uint `uri:"itemID" binding:"required"`
}
