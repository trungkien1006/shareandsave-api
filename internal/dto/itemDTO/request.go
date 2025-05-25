package itemDTO

type CreateItemRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type CreateItemResponse struct {
	Item ItemDTO `json:"item"`
}

type CreateItemResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    CreateItemResponse `json:"data"`
}

type UpdateItemRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
