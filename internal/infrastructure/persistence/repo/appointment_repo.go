package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/appointment"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type AppointmentRepoDB struct {
	db *gorm.DB
}

func NewAppointmentRepoDB(db *gorm.DB) *AppointmentRepoDB {
	return &AppointmentRepoDB{db: db}
}

func (r *AppointmentRepoDB) Create(ctx context.Context, appointments map[uint]appointment.Appointment) error {
	var (
		dbAppointments []dbmodel.Appointment
	)

	tx := r.db.Debug().WithContext(ctx).Begin()

	for _, value := range appointments {
		dbAppointments = append(dbAppointments, dbmodel.AppointmentDomainToDB(value))
	}

	if err := tx.Model(&dbmodel.Appointment{}).Create(&dbAppointments).Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi khởi tạo phiếu hẹn: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi commit transaction: " + err.Error())
	}

	return nil
}
