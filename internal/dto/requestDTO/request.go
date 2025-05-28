package requestdto

import (
	"final_project/internal/pkg/enums"
	"time"
)

type CreateRequestSendOldItem struct {
	ID                  uint              `json:"id" binding:"required"`
	Email               string            `json:"email" example:"john@gmail.com"`
	FullName            string            `json:"fullName" example:"John Doe"`
	PhoneNumber         string            `json:"phoneNumber" example:"0123456789"`
	UserID              uint              `json:"userId"`
	RequestType         enums.RequestType `json:"requestType" binding:"required;oneof=1 2 3" example:"1"` // 1: Send Old Item, 2: Request Item, 3: Request Post, 4: Reply Post
	Description         string            `json:"description" example:"I want to send this old item to the charity organization."`
	IsAnonymous         bool              `json:"isAnonymous" binding:"oneof=true false"`                                      // true: anonymous, false: not anonymous
	AppointmentTime     time.Time         `json:"appointmentTime" binding:"required" example:"2025-05-28T07:23:45Z"`           // Time in RFC3339 format
	AppointmentLocation string            `json:"appointmentLocation" binding:"required" example:"123 Main St, City, Country"` // Location of the appointment
}
