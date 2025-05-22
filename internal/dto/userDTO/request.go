package userDTO

type GetUserRequest struct {
	Page   int    `query:"page" binding:"default:1"`
	Limit  int    `query:"limit" binding:"default:8"`
	Sort   string `query:"sort" binding:"default:'ASC';oneof:ASC DESC"`
	Order  string `query:"order" binding:"default:''"`
	Filter string `query:"filter" binding:"default:''"`
}

type GetUserByIDRequest struct {
	UserID int `query:"userID" binding:"required"`
}
