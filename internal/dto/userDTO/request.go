package userDTO

import "final_project/internal/pkg/enums"

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

type CreateUserRequest struct {
	FullName    string           `json:"fullName" binding:"required"`
	Email       string           `json:"email" binding:"required"`
	PhoneNumber string           `json:"phoneNumber"`
	Password    string           `json:"password" binding:"required,min=8"`
	Status      enums.UserStatus `json:"status" binding:"oneof=0 1 2"`
	Address     string           `json:"address"`
	GoodPoint   int              `json:"goodPoint"`
}

type UpdateUserRequest struct {
	ID          int              `json:"id"`
	FullName    string           `json:"fullName"`
	Email       string           `json:"email"`
	PhoneNumber string           `json:"phoneNumber"`
	Password    string           `json:"password" binding:"min=8"`
	Status      enums.UserStatus `json:"status" binding:"oneof=0 1 2"`
	Address     string           `json:"address"`
	GoodPoint   int              `json:"goodPoint"`
}

func (r *GetUserRequest) SetDefaults() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
}
