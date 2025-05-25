package itemDTO

// Request dùng cho lấy 1 item theo ID
type GetItemByIDRequest struct {
	ItemID uint `uri:"itemID" binding:"required"`
}

// Response trả về một item
type GetItemByIDResponse struct {
	Item ItemDTO `json:"item"`
}

// Wrapper cho response lấy 1 item (thêm code/message)
type GetItemByIDResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    GetItemByIDResponse `json:"data"`
}

// Request dùng để lấy danh sách item
type GetAllItemRequest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`
	Filter string `form:"filter"` // JSON string nếu có
}

func (r *GetAllItemRequest) SetDefaults() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.Limit <= 0 {
		r.Limit = 10
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
}

// Response trả về danh sách item
type GetItemResponse struct {
	Items     []ItemDTO `json:"items"`
	TotalPage int       `json:"totalPage"`
}

// Wrapper cho response danh sách item
type GetItemResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    GetItemResponse `json:"data"`
}

// Response trả về item sau khi update
type UpdateItemResponse struct {
	Item ItemDTO `json:"item"`
}

// Wrapper cho response update item
type UpdateItemResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    UpdateItemResponse `json:"data"`
}

// Request dùng để delete item theo ID
type DeleteItemRequest struct {
	ItemID uint `uri:"itemID" binding:"required"`
}

// Response trả về sau khi delete (không có dữ liệu)
type DeleteItemResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Wrapper cho response delete item
type DeleteItemResponseWrapper struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
