package admindto

import "final_project/internal/pkg/enums"

type GetAllAdminRequest struct {
	Page   int    `query:"page" binding:"omitempty,min=1"`
	Limit  int    `query:"limit" binding:"omitempty,min=8"`
	Sort   string `query:"sort" binding:"omitempty,oneof=ASC DESC"`
	Order  string `query:"order"`
	Filter string `query:"filter"`
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
