package appointment

import "context"

type Repository interface {
	Create(ctx context.Context, appointments map[uint]Appointment) error
}
