package request

import (
	"final_project/internal/pkg/enums"
	"time"
)

type SendRequest struct {
	ID                  uint
	UserID              uint
	Type                enums.RequestType // 1: Send Old Item, 2: Request Item, 3: Request Post, 4: Reply Post
	Description         string
	Status              int8
	ReplyMessage        string
	AppointmentTime     time.Time
	AppointmentLocation string
}

type ReceiveRequest struct {
	ID                  uint
	UserID              uint
	Type                enums.RequestType // 1: Send Old Item, 2: Request Item, 3: Request Post, 4: Reply Post
	Description         string
	Status              int8
	ItemWarehouseID     *uint
	PostID              *uint
	ReplyMessage        string
	AppointmentTime     time.Time
	AppointmentLocation string
}
