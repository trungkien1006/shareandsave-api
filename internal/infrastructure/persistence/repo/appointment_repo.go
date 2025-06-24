package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/appointment"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type AppointmentRepoDB struct {
	db *gorm.DB
}

func NewAppointmentRepoDB(db *gorm.DB) *AppointmentRepoDB {
	return &AppointmentRepoDB{db: db}
}

func (r *AppointmentRepoDB) GetAll(ctx context.Context, appointments *[]appointment.Appointment, req appointment.FilterAllAppointment, userID uint) (int, error) {
	var (
		query          *gorm.DB
		totalRecords   int64
		dbAppointments []dbmodel.Appointment
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Appointment{}).
		Table("appointment as a").
		Joins("JOIN user ON user.id = a.user_id").
		Preload("User").
		Preload("AppointmentItem").
		Preload("AppointmentItem.Item").
		Preload("AppointmentItem.Item.Category")

	if userID != 0 {
		query.Where("a.user_id = ? ", userID)
	}

	if req.SearchBy != "" && req.SearchValue != "" {
		column := strcase.ToSnake(req.SearchBy) // "fullName" -> "full_name"

		if column == "user_name" {
			column = "user.full_name"
		} else {
			column = "a." + column
		}

		query.Where(column+" LIKE ? ", "%"+req.SearchValue+"%")
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, errors.New("Có lỗi khi truy vấn danh sách phiếu hẹn: " + err.Error())
	}

	if req.Sort != "" && req.Order != "" {
		query = query.Order("a." + strcase.ToSnake(req.Sort) + " " + req.Order)
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&dbAppointments).Error; err != nil {
		return 0, errors.New("Có lỗi khi truy vấn danh sách phiếu hẹn: " + err.Error())
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	for _, value := range dbAppointments {
		*appointments = append(*appointments, dbmodel.AppointmentDBToDomain(value))
	}

	return totalPages, nil
}

func (r *AppointmentRepoDB) GetByID(ctx context.Context, appointment *appointment.Appointment, appointmentID uint) error {
	var dbAppointment dbmodel.Appointment

	if err := r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Appointment{}).
		Table("appointment as a").
		Joins("JOIN user ON user.id = a.user_id").
		Where("a.id = ? ", appointmentID).
		Preload("User").
		Preload("AppointmentItem").
		Preload("AppointmentItem.Item").
		Preload("AppointmentItem.Item.Category").
		First(&dbAppointment).Error; err != nil {
		return errors.New("Có lỗi khi truy vấn phiếu hẹn bằng id: " + err.Error())
	}

	*appointment = dbmodel.AppointmentDBToDomain(dbAppointment)

	return nil
}

func (r *AppointmentRepoDB) CreateBatch(ctx context.Context, appointments map[uint]appointment.Appointment) error {
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
