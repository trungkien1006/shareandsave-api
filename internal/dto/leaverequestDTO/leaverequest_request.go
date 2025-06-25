package leaverequestdto

import (
	"final_project/internal/pkg/enums"
	"time"
)

type GetAllLeaveRequestRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort" binding:"omitempty,oneof=startDate endDate totalDays createdAt"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy" binding:"omitempty,oneof=leaveType reason userName"`
	SearchValue string `form:"searchValue"`
}

func (r *GetAllLeaveRequestRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
}

type CreateLeaveRequestRequest struct {
	ID        int                    `json:"id" binding:"required"`
	UserID    int                    `json:"userID" binding:"required"`
	LeaveType enums.LeaveRequestType `json:"leaveType" binding:"required"`
	Reason    string                 `json:"reason" binding:"required"`
	StartDate time.Time              `json:"startDate" binding:"required"`
	EndDate   time.Time              `json:"endDate" binding:"required"`
}
