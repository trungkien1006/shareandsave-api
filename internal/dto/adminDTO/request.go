package admindto

import "final_project/internal/pkg/enums"

type CreateAdminRequest struct {
	Fullname string           `json:"fullname" binding:"required"`
	Email    string           `json:"email" binding:"required"`
	Password string           `json:"password" binding:"required,min=8"`
	Status   enums.UserStatus `json:"status" binding:"oneof=0 1 2"`
	RoleID   uint             `json:"roleId" binding:"required"`
}

type UpdateAdminRequest struct {
	ID       uint             `json:"id" binding:"required"`
	Fullname string           `json:"fullname"`
	Password string           `json:"password"`
	Status   enums.UserStatus `json:"status" binding:"oneof=0 1 2"`
	RoleID   uint             `json:"roleId"`
}

type DeleteAdminRequest struct {
	AdminID int `uri:"adminID" binding:"required"`
}
