package userdto

import "final_project/internal/pkg/enums"

type GetUserRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy"`
	SearchValue string `form:"searchValue"`
}

func (r *GetUserRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
	if r.Sort == "" {
		r.Sort = "id"
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
}

type GetUserByIDRequest struct {
	UserID int `uri:"userID" binding:"required"`
}

type CreateUserRequest struct {
	FullName    string           `json:"fullName" binding:"required" example:"John Doe"`
	Email       string           `json:"email" binding:"required" example:"john@gmail.com"`
	PhoneNumber string           `json:"phoneNumber" example:"0123456789"`
	Password    string           `json:"password" binding:"required,min=8"`
	Status      enums.UserStatus `json:"status" binding:"oneof=0 1 2" example:"1"` // 0: inactive, 1: active, 2: banned
	Address     string           `json:"address" exapmple:"123 Main St, City, Country"`
	Avatar      string           `json:"avatar"`
	GoodPoint   int              `json:"goodPoint"`
}

type UpdateUserRequest struct {
	FullName    string           `json:"fullName" example:"John Doe"`
	PhoneNumber string           `json:"phoneNumber" example:"0123456789"`
	Avatar      string           `json:"avatar"`
	Status      enums.UserStatus `json:"status" binding:"oneof=0 1 2" example:"1"` // 0: inactive, 1: active, 2: banned
	Address     string           `json:"address" exapmple:"123 Main St, City, Country"`
	GoodPoint   int              `json:"goodPoint"`
	Major       string           `json:"major" example:"Information Technology"`
}

type DeleteUserRequest struct {
	UserID int `uri:"userID" binding:"required"`
}
