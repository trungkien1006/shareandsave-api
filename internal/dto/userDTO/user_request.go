package userdto

import "final_project/internal/pkg/enums"

type GetClientRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort" binding:"omitempty,oneof=createdAt goodPoint"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy" binding:"omitempty,oneof=fullName email phoneNumber status"`
	SearchValue string `form:"searchValue"`
}

func (r *GetClientRequest) SetDefault() {
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

type GetAdminRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort" binding:"omitempty,oneof=createdAt goodPoint"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy" binding:"omitempty,oneof=fullName email phoneNumber status roleName"`
	SearchValue string `form:"searchValue"`
}

func (r *GetAdminRequest) SetDefault() {
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

// Get client by id
type GetClientByIDRequest struct {
	ClientID int `uri:"clientID" binding:"required"`
}

// Get admin by id
type GetAdminByIDRequest struct {
	AdminID int `uri:"adminID" binding:"required"`
}

type CreateClientRequest struct {
	FullName    string           `json:"fullName" binding:"required" example:"John Doe"`
	Email       string           `json:"email" binding:"required" example:"john@gmail.com"`
	PhoneNumber string           `json:"phoneNumber" example:"0123456789"`
	Password    string           `json:"password" binding:"required,min=8"`
	Status      enums.UserStatus `json:"status" binding:"oneof=0 1 2" example:"1"` // 0: inactive, 1: active, 2: banned
	Address     string           `json:"address" exapmple:"123 Main St, City, Country"`
	Avatar      string           `json:"avatar"`
	GoodPoint   int              `json:"goodPoint"`
}

type CreateAdminRequest struct {
	RoleID      uint             `json:"roleID" binding:"required"`
	FullName    string           `json:"fullName" binding:"required" example:"John Doe"`
	Email       string           `json:"email" binding:"required" example:"john@gmail.com"`
	PhoneNumber string           `json:"phoneNumber" example:"0123456789"`
	Password    string           `json:"password" binding:"required,min=8"`
	Status      enums.UserStatus `json:"status" binding:"oneof=0 1 2" example:"1"` // 0: inactive, 1: active, 2: banned
	Address     string           `json:"address" exapmple:"123 Main St, City, Country"`
	Avatar      string           `json:"avatar"`
	GoodPoint   int              `json:"goodPoint"`
}

type UpdateClientRequest struct {
	FullName    string           `json:"fullName" example:"John Doe"`
	PhoneNumber string           `json:"phoneNumber" example:"0123456789"`
	Avatar      string           `json:"avatar"`
	Status      enums.UserStatus `json:"status" binding:"oneof=0 1 2" example:"1"` // 0: inactive, 1: active, 2: banned
	Address     string           `json:"address" exapmple:"123 Main St, City, Country"`
	GoodPoint   int              `json:"goodPoint"`
	Major       string           `json:"major" example:"Information Technology"`
}

type UpdateAdminRequest struct {
	RoleID      uint             `json:"roleID"`
	FullName    string           `json:"fullName" example:"John Doe"`
	PhoneNumber string           `json:"phoneNumber" example:"0123456789"`
	Avatar      string           `json:"avatar"`
	Status      enums.UserStatus `json:"status" binding:"oneof=0 1 2" example:"1"` // 0: inactive, 1: active, 2: banned
	Address     string           `json:"address" exapmple:"123 Main St, City, Country"`
	GoodPoint   int              `json:"goodPoint"`
	Major       string           `json:"major" example:"Information Technology"`
}

type DeleteClientRequest struct {
	CLientID int `uri:"clientID" binding:"required"`
}

type DeleteAdminRequest struct {
	AdminID int `uri:"adminID" binding:"required"`
}
