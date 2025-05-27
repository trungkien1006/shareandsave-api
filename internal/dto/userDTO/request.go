package userDTO

import "final_project/internal/pkg/enums"

type GetUserRequest struct {
	Page   int    `query:"page" binding:"min=1"`
	Limit  int    `query:"limit" binding:"min=8"`
	Sort   string `query:"sort" binding:"omitempty,oneof=ASC DESC"`
	Order  string `query:"order"`
	Filter string `query:"filter"`
}

type GetUserByIDRequest struct {
	UserID int `uri:"userID" binding:"required"`
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
	ID          uint             `json:"id"`
	FullName    string           `json:"fullName"`
	PhoneNumber string           `json:"phoneNumber"`
	Avatar      string           `json:"avatar"`
	Status      enums.UserStatus `json:"status" binding:"oneof=0 1 2"`
	Address     string           `json:"address"`
	GoodPoint   int              `json:"goodPoint"`
	Major       string           `json:"major"`
}

type DeleteUserRequest struct {
	UserID int `uri:"userID" binding:"required"`
}
