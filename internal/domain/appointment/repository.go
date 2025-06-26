package appointment

import (
	"context"
	"time"
)

type Repository interface {
	CreateBatch(ctx context.Context, appointments map[uint]Appointment) error
	GetAll(ctx context.Context, appointments *[]Appointment, req FilterAllAppointment, userID uint) (int, error)
	GetByID(ctx context.Context, appointment *Appointment, appointmentID uint) error
	Update(ctx context.Context, domainAppointment Appointment) error
	UpdateBatch(ctx context.Context, appointments []Appointment) error
	IsInDay(ctx context.Context, day time.Time) (bool, error)
}
