package appointmentapp

import (
	"context"
	"errors"
	"final_project/internal/domain/appointment"
	"time"
)

type UseCase struct {
	appointmentRepo appointment.Repository
}

func NewUseCase(appointmentRepo appointment.Repository) *UseCase {
	return &UseCase{appointmentRepo: appointmentRepo}
}

func (uc *UseCase) GetAll(ctx context.Context, appointments *[]appointment.Appointment, filter appointment.FilterAllAppointment, userID uint) (int, error) {
	totalPage, err := uc.appointmentRepo.GetAll(ctx, appointments, filter, userID)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetByID(ctx context.Context, appointment *appointment.Appointment, appointmentID uint) error {
	if err := uc.appointmentRepo.GetByID(ctx, appointment, appointmentID); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) Update(ctx context.Context, domainAppointment appointment.Appointment, appointmentID uint) error {
	var updateAppointment appointment.Appointment

	if err := uc.appointmentRepo.GetByID(ctx, &updateAppointment, appointmentID); err != nil {
		return err
	}

	if domainAppointment.Status != 0 {
		updateAppointment.Status = domainAppointment.Status
	}

	if domainAppointment.EndTime.Before(domainAppointment.StartTime) {
		return errors.New("Thời gian kết thúc không được < thời gian bắt đầu")
	}

	if domainAppointment.StartTime.Sub(domainAppointment.EndTime) > time.Hour {
		return errors.New("Thời gian kết thúc và bắt đầu chỉ được cách nhau 1 tiếng")
	}

	if !domainAppointment.StartTime.IsZero() && !domainAppointment.EndTime.IsZero() {
		updateAppointment.StartTime = domainAppointment.StartTime
		updateAppointment.EndTime = domainAppointment.EndTime
	} else if !domainAppointment.StartTime.IsZero() || !domainAppointment.EndTime.IsZero() {
		return errors.New("Muốn cập nhật thời gian phải gửi cả thời gian bắt đầu và kết thúc")
	}

	updateAppointment.ID = appointmentID

	if err := uc.appointmentRepo.Update(ctx, updateAppointment); err != nil {
		return err
	}

	return nil
}
