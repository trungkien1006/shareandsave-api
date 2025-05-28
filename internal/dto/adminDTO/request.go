package admindto

import "final_project/internal/pkg/enums"

type GetAllAdminRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy"`
	SearchValue string `form:"searchValue"`
}

func (r *GetAllAdminRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 10
	}
	if r.Sort == "" {
		r.Sort = "id"
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
}

type GetAdminByIDRequest struct {
	AdminID int `uri:"adminID" binding:"required"`
}

type CreateAdminRequest struct {
	FullName string           `json:"fullName" binding:"required" example:"John Doe"`
	Email    string           `json:"email" binding:"required" example:"john@gmail.com"`
	Password string           `json:"password" binding:"required,min=8"`
	Status   enums.UserStatus `json:"status" binding:"oneof=0 1 2" example:"0"` // 0: Inactive, 1: Active, 2: Suspended
	RoleID   uint             `json:"roleId" binding:"required" example:"1"`    // Role ID must be provided, e.g., 1 for Admin, 2 for Moderator
}

type UpdateAdminRequest struct {
	ID       uint             `json:"id" binding:"required"`
	FullName string           `json:"fullName" example:"John Doe"`
	Password string           `json:"password"`
	Status   enums.UserStatus `json:"status" binding:"oneof=0 1 2" example:"0"` // 0: Inactive, 1: Active, 2: Suspended
	RoleID   uint             `json:"roleId" example:"1"`                       // Role ID must be provided, e.g., 1 for Admin, 2 for Moderator
}

type DeleteAdminRequest struct {
	AdminID int `uri:"adminID" binding:"required"`
}
