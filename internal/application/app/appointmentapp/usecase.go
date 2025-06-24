package appointmentapp

import (
	"context"
	"final_project/internal/domain/appointment"
)

type UseCase struct {
	appointmentRepo appointment.Repository
}

func NewUseCase(appointmentRepo appointment.Repository) *UseCase {
	return &UseCase{appointmentRepo: appointmentRepo}
}

func (uc *UseCase) GetAll(ctx context.Context, appointments *[]appointment.Appointment, filter appointment.FilterAllAppointment) (int, error) {
	totalPage, err := uc.appointmentRepo.GetAll(ctx, appointments, filter)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}
