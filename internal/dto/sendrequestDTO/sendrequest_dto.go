package sendrequestdto

import (
	sendrequest "final_project/internal/domain/send_request"
	"final_project/internal/pkg/enums"
	"time"
)

type RequestSendOldItem struct {
	ID                  uint              `json:"id"`
	UserID              uint              `json:"userId"`
	Type                enums.RequestType `json:"type"` // 1: Send Old Item, 2: Request Item, 3: Request Post, 4: Reply Post
	Description         string            `json:"description"`
	AppointmentTime     time.Time         `json:"appointmentTime"`     // Time in RFC3339 format
	AppointmentLocation string            `json:"appointmentLocation"` // Location of the appointment
}

func ToRequestDTO(u sendrequest.SendRequest) RequestSendOldItem {
	return RequestSendOldItem{
		ID:                  u.ID,
		UserID:              u.UserID,
		Type:                enums.RequestType(u.Type),
		Description:         u.Description,
		AppointmentTime:     u.AppointmentTime,
		AppointmentLocation: u.AppointmentLocation,
	}
}
