package request

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

type ReceiveRequest struct {
	ID                  uint
	UserID              uint
	Type                int
	Description         string
	Status              int8
	ItemWarehouseID     *uint
	PostID              *uint
	ReplyMessage        string
	AppointmentTime     time.Time
	AppointmentLocation string
}
