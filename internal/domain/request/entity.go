package request

import "time"

type Request struct {
	ID                  uint
	UserID              uint
	RequestType         int
	Description         string
	IsAnonymous         bool
	Status              int8
	ItemWarehouseID     *uint
	PostID              *uint
	ReplyMessage        string
	AppointmentTime     time.Time
	AppointmentLocation string
}
