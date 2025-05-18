package userDTO

type GetUserRequest struct {
	Page   int    `json:"page" binding:"required"`
	Limit  int    `json:"limit" binding:"required"`
	Sort   string `json:"sort" binding:"required; oneof:ASC DESC"`
	Order  string `json:"order" binding:"required"`
	Filter string `json:"filter" binding:"required"`
}
