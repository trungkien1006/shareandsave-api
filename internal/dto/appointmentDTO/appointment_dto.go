package appointmentdto

import "time"

type AppointmentDTO struct {
	ID               uint                 `json:"id"`
	UserID           uint                 `json:"userID"`
	StartTime        time.Time            `json:"startTime"`
	EndTime          time.Time            `json:"endTime"`
	Status           int                  `json:"status"`
	AppointmentItems []AppointmentItemDTO `json:"appointmentItems"`
	CreatedAt        time.Time            `json:"createdAt"`
}

type AppointmentItemDTO struct {
	ID              uint `json:"id"`
	AppointmentID   uint `json:"appointmentID"`
	ItemID          uint `json:"itemID"`
	ActualQuantity  int  `json:"actualQuantity"`
	MissingQuantity int  `json:"missingQuantity"`
}
