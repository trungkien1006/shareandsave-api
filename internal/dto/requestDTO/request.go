package requestdto

import "time"

type CreateRequestSendOldItem struct {
	ID                  uint      `json:"id" binding:"required"`
	Email               string    `json:"email"`
	FullName            string    `json:"fullName"`
	PhoneNumber         string    `json:"phoneNumber"`
	UserID              uint      `json:"userId"`
	RequestType         int       `json:"requestType" binding:"required;oneof=SendOldItem SendLoseItem ReceiveOldItem"` // 1: Send Old Item, 2: Request Item, 3: Request Post, 4: Reply Post
	Description         string    `json:"description"`
	IsAnonymous         bool      `json:"isAnonymous" binding:"oneof=true false"` // true: anonymous, false: not anonymous
	AppointmentTime     time.Time `json:"appointmentTime" binding:"required"`     // Time in RFC3339 format
	AppointmentLocation string    `json:"appointmentLocation" binding:"required"` // Location of the appointment
}
