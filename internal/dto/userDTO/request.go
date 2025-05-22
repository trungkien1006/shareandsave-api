package userDTO

type GetUserRequest struct {
	Page   int    `query:"page" binding:"min:1"`
	Limit  int    `query:"limit" binding:"min:1"`
	Sort   string `query:"sort" binding:"oneof:ASC DESC"`
	Order  string `query:"order"`
	Filter string `query:"filter"`
}

type GetUserByIDRequest struct {
	UserID int `query:"userID" binding:"required"`
}
