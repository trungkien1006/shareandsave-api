package appointmentdto

import (
	"final_project/internal/pkg/enums"
	"time"
)

type GetAllAppointmentRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort" binding:"omitempty,oneof=startTime endTime"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy" binding:"omitempty,oneof=status userName"`
	SearchValue string `form:"searchValue"`
}

func (r *GetAllAppointmentRequest) SetDefault() {
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

type GetAllMyAppointmentRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort" binding:"omitempty,oneof=startTime endTime"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy" binding:"omitempty,oneof=status"`
	SearchValue string `form:"searchValue"`
}

func (r *GetAllMyAppointmentRequest) SetDefault() {
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

type GetAppointmentByIDRequest struct {
	AppointmentID uint `uri:"appointmentID"`
}

type UpdateAppointmentRequest struct {
	StartTime time.Time               `json:"startTime"`
	EndTime   time.Time               `json:"endTime"`
	Status    enums.AppointmentStatus `json:"status"`
}

type UpdateBatchAppointmentRequest struct {
	AppointmentIDs []uint                     `json:"appointmentIDs" binding:"required"`
	Appointments   []UpdateAppointmentRequest `json:"appointments" binding:"required"`
}
