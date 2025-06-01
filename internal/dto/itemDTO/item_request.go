package itemdto

type GetAllItemRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy" binding:"omitempty,oneof=categoryID"`
	SearchValue string `form:"searchValue"`
}

func (r *GetAllItemRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 10
	}
	if r.Sort == "" {
		r.Sort = "id"
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
}

// Request dùng cho lấy 1 item theo ID
type GetItemByIDRequest struct {
	ItemID uint `uri:"itemID" binding:"required"`
}

type CreateItemRequest struct {
	CategoryID  uint   `json:"categoryID" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type UpdateItemRequest struct {
	CategoryID  uint   `json:"categoryID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// Request dùng để delete item theo ID
type DeleteItemRequest struct {
	ItemID uint `uri:"itemID" binding:"required"`
}
