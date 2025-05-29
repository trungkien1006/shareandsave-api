package sendrequestdto

import (
	"time"
)

type CreateRequestSendOldItem struct {
	Email               string    `json:"email" example:"john@gmail.com"`
	FullName            string    `json:"fullName" example:"John Doe"`
	PhoneNumber         string    `json:"phoneNumber" example:"0123456789"`
	UserID              uint      `json:"userId"`
	Description         string    `json:"description" example:"I want to send this old item to the charity organization."`
	IsAnonymous         bool      `json:"isAnonymous"`                                                                 // true: anonymous, false: not anonymous
	AppointmentTime     time.Time `json:"appointmentTime" binding:"required" example:"2025-05-28T07:23:45Z"`           // Time in RFC3339 format
	AppointmentLocation string    `json:"appointmentLocation" binding:"required" example:"123 Main St, City, Country"` // Location of the appointment
}
