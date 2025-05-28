package requestdto

import (
	"time"
)

type RequestSendOldItem struct {
	ID                  uint      `json:"id"`
	UserID              uint      `json:"userId"`
	RequestType         int       `json:"requestType"` // 1: Send Old Item, 2: Request Item, 3: Request Post, 4: Reply Post
	Description         string    `json:"description"`
	IsAnonymous         bool      `json:"isAnonymous"`         // true: anonymous, false: not anonymous
	AppointmentTime     time.Time `json:"appointmentTime"`     // Time in RFC3339 format
	AppointmentLocation string    `json:"appointmentLocation"` // Location of the appointment
}
