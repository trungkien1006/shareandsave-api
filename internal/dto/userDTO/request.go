package userDTO

type GetUserRequest struct {
	Page   int    `query:"page" binding:"omitempty,min=1"`
	Limit  int    `query:"limit" binding:"omitempty,min=1"`
	Sort   string `query:"sort" binding:"omitempty,oneof=ASC DESC"`
	Order  string `query:"order"`
	Filter string `query:"filter"`
}

type GetUserByIDRequest struct {
	UserID int `query:"userID" binding:"required"`
}

func (r *GetUserRequest) SetDefaults() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
}
