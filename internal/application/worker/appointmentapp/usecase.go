package appointmentapp

import "final_project/internal/domain/appointment"

type UseCase struct {
	appointmentRepo appointment.Repository
}

func NewUseCase(appointmentRepo appointment.Repository) *UseCase {
	return &UseCase{appointmentRepo: appointmentRepo}
}
