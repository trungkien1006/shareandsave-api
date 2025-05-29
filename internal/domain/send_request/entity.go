package sendrequest

import "time"

type SendRequest struct {
	ID                  uint
	UserID              uint
	Type                int
	Description         string
	Status              int8
	ReplyMessage        string
	AppointmentTime     time.Time
	AppointmentLocation string
}
