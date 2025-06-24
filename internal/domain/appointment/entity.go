package appointment

import "time"

type Appointment struct {
	ID               uint
	UserID           uint
	StartTime        time.Time
	EndTime          time.Time
	Status           int
	AppointmentItems []AppointmentItem
}

type AppointmentItem struct {
	ID              uint
	AppointmentID   uint
	ItemID          uint
	ActualQuantity  int
	MissingQuantity int
}
