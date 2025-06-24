package appointment

import "context"

type Repository interface {
	CreateBatch(ctx context.Context, appointments map[uint]Appointment) error
	GetAll(ctx context.Context, appointments *[]Appointment, req FilterAllAppointment, userID uint) (int, error)
}
