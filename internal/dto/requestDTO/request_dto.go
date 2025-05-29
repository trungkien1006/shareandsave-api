package requestdto

import (
	"final_project/internal/pkg/enums"
	"time"
)

type RequestSendOldItem struct {
	ID                  uint                `json:"id"`
	UserID              uint                `json:"userId"`
	Type                enums.RequestType   `json:"type"` // 1: Send Old Item, 2: Request Item, 3: Request Post, 4: Reply Post
	Description         string              `json:"description"`
	ReplyMessage        string              `json:"replyMessage,omitempty"` // Optional, used for replies
	AppointmentTime     time.Time           `json:"appointmentTime"`        // Time in RFC3339 format
	AppointmentLocation string              `json:"appointmentLocation"`    // Location of the appointment
	Status              enums.RequestStatus `json:"status"`                 // 0: Pending, 1: Accepted, 2: Rejected, 3: Completed
	IsAnonymous         bool                `json:"isAnonymous"`
}
