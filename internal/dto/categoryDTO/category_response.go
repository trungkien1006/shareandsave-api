package categorydto

// Response trả về danh sách item
type GetCategoryResponse struct {
	Categories []CategoryDTO `json:"categories"`
}

// Wrapper cho response danh sách Category
type GetCategoryResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    GetCategoryResponse `json:"data"`
}
