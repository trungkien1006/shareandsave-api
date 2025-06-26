package appointmentapp

import (
	"context"
	"errors"
	"final_project/internal/domain/appointment"
	"strconv"
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

	// Kiểm tra nếu thời gian kết thúc < thời gian bắt đầu
	if domainAppointment.EndTime.Before(domainAppointment.StartTime) {
		return errors.New("Thời gian kết thúc không được < thời gian bắt đầu")
	}

	// Kiểm tra thời gian cách nhau không quá 1 tiếng
	if domainAppointment.StartTime.Sub(domainAppointment.EndTime) > time.Hour {
		return errors.New("Thời gian kết thúc và bắt đầu chỉ được cách nhau 1 tiếng")
	}

	// Kiểm tra nếu cập nhật thời gian thì phải gửi cả hai
	if !domainAppointment.StartTime.IsZero() && !domainAppointment.EndTime.IsZero() {
		// 🆕 Kiểm tra thời gian không được ở quá khứ
		now := time.Now()
		if domainAppointment.StartTime.Before(now) || domainAppointment.EndTime.Before(now) {
			return errors.New("Thời gian bắt đầu hoặc kết thúc không được ở quá khứ")
		}

		updateAppointment.StartTime = domainAppointment.StartTime
		updateAppointment.EndTime = domainAppointment.EndTime
	} else if !domainAppointment.StartTime.IsZero() || !domainAppointment.EndTime.IsZero() {
		return errors.New("Muốn cập nhật thời gian phải gửi cả thời gian bắt đầu và kết thúc")
	}

	if domainAppointment.Status != 0 {
		updateAppointment.Status = domainAppointment.Status
	}

	updateAppointment.ID = appointmentID

	if err := uc.appointmentRepo.Update(ctx, updateAppointment); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdateBatch(ctx context.Context, domainAppointment []appointment.Appointment, appointmentID []uint) error {
	updateAppointment := make([]appointment.Appointment, 0)

	if len(domainAppointment) != len(appointmentID) {
		return errors.New("Số lượng appointment và ID không khớp: " + strconv.Itoa(len(domainAppointment)) + " != " + strconv.Itoa(len(appointmentID)))
	}

	for key, value := range domainAppointment {
		if err := uc.appointmentRepo.GetByID(ctx, &updateAppointment[key], appointmentID[key]); err != nil {
			return err
		}

		// Kiểm tra nếu thời gian kết thúc < thời gian bắt đầu
		if value.EndTime.Before(value.StartTime) {
			return errors.New("Thời gian kết thúc không được < thời gian bắt đầu")
		}

		// Kiểm tra thời gian cách nhau không quá 1 tiếng
		if value.StartTime.Sub(value.EndTime) > time.Hour {
			return errors.New("Thời gian kết thúc và bắt đầu chỉ được cách nhau 1 tiếng")
		}

		// Kiểm tra nếu cập nhật thời gian thì phải gửi cả hai
		if !value.StartTime.IsZero() && !value.EndTime.IsZero() {
			// 🆕 Kiểm tra thời gian không được ở quá khứ
			now := time.Now()
			if value.StartTime.Before(now) || value.EndTime.Before(now) {
				return errors.New("Thời gian bắt đầu hoặc kết thúc không được ở quá khứ")
			}

			updateAppointment[key].StartTime = value.StartTime
			updateAppointment[key].EndTime = value.EndTime
		} else if !value.StartTime.IsZero() || !value.EndTime.IsZero() {
			return errors.New("Muốn cập nhật thời gian phải gửi cả thời gian bắt đầu và kết thúc")
		}

		if value.Status != 0 {
			updateAppointment[key].Status = value.Status
		}

		updateAppointment[key].ID = appointmentID[key]
	}

	if err := uc.appointmentRepo.UpdateBatch(ctx, updateAppointment); err != nil {
		return err
	}

	return nil
}
