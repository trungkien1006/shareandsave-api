package admindto

import "final_project/internal/pkg/enums"

type GetAllAdminRequest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Sort   string `form:"sort" binding:"omitempty,oneof=ASC DESC"`
	Order  string `form:"order"`
	Filter string `form:"filter"`
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
	FullName string           `json:"fullName" binding:"required"`
	Email    string           `json:"email" binding:"required"`
	Password string           `json:"password" binding:"required,min=8"`
	Status   enums.UserStatus `json:"status" binding:"oneof=0 1 2"`
	RoleID   uint             `json:"roleId" binding:"required"`
}

type UpdateAdminRequest struct {
	ID       uint             `json:"id" binding:"required"`
	FullName string           `json:"fullName"`
	Password string           `json:"password"`
	Status   enums.UserStatus `json:"status" binding:"oneof=0 1 2"`
	RoleID   uint             `json:"roleId"`
}

type DeleteAdminRequest struct {
	AdminID int `uri:"adminID" binding:"required"`
}
