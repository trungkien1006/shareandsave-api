package appointment

import "time"

type Appointment struct {
	ID               uint
	UserID           uint
	StartTime        time.Time
	EndTime          time.Time
	Status           int
	AppointmentItems []AppointmentItem
	CreatedAt        time.Time
}

type AppointmentItem struct {
	ID              uint
	AppointmentID   uint
	ItemID          uint
	ActualQuantity  int
	MissingQuantity int
}

type FilterAllAppointment struct {
	Page        int
	Limit       int
	Sort        string
	Order       string
	SearchBy    string
	SearchValue string
}
