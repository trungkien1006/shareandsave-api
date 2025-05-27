package itemDTO

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

type CreateItemResponse struct {
	Item ItemDTO `json:"item"`
}

type CreateItemResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    CreateItemResponse `json:"data"`
}
