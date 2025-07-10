package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/appointment"
	"final_project/internal/domain/notification"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"fmt"
	"time"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type AppointmentRepoDB struct {
	db          *gorm.DB
	notiService notification.Service
}

func NewAppointmentRepoDB(db *gorm.DB, notiService notification.Service) *AppointmentRepoDB {
	return &AppointmentRepoDB{db: db, notiService: notiService}
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

func (r *AppointmentRepoDB) Update(ctx context.Context, domainAppointment appointment.Appointment) error {
	var dbAppointment dbmodel.Appointment

	dbAppointment = dbmodel.AppointmentDomainToDB(domainAppointment)

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Appointment{}).
		Where("id = ?", dbAppointment.ID).
		Updates(&dbAppointment).Error; err != nil {
		return errors.New("Có lỗi khi cập nhật phiếu hẹn: " + err.Error())
	}

	appointmentDay := fmt.Sprintf("%02d:%02d %02d/%02d/%04d",
		domainAppointment.StartTime.Hour(),
		domainAppointment.StartTime.Minute(),
		domainAppointment.StartTime.Day(),
		int(domainAppointment.StartTime.Month()),
		domainAppointment.StartTime.Year(),
	)

	noti := notification.Notification{
		Type:       "system",
		TargetType: "appointment",
		TargetID:   dbAppointment.ID,
		IsRead:     false,
		SenderID:   nil,
		ReceiverID: &dbAppointment.UserID,
		Content:    "Có 1 cuộc hẹn của bạn được hẹn lại vào lúc: " + appointmentDay,
	}

	if err := r.notiService.CreateAndPushSocket(ctx, &noti); err != nil {
		return errors.New("Có lỗi khi thêm thông báo: " + err.Error())
	}

	return nil
}

// UpdateBatch cập nhật nhiều appointment cùng lúc (theo ID)
func (r *AppointmentRepoDB) UpdateBatch(ctx context.Context, appointments []appointment.Appointment) error {
	tx := r.db.WithContext(ctx).Begin()

	for _, domainAppointment := range appointments {
		dbAppointment := dbmodel.AppointmentDomainToDB(domainAppointment)
		if err := tx.Model(&dbmodel.Appointment{}).
			Where("id = ?", dbAppointment.ID).
			Updates(&dbAppointment).Error; err != nil {
			tx.Rollback()
			return errors.New("Có lỗi khi cập nhật phiếu hẹn: " + err.Error())
		}
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("Có lỗi khi commit transaction: " + err.Error())
	}

	for _, value := range appointments {
		appointmentDay := fmt.Sprintf("%02d:%02d %02d/%02d/%04d",
			value.StartTime.Hour(),
			value.StartTime.Minute(),
			value.StartTime.Day(),
			int(value.StartTime.Month()),
			value.StartTime.Year(),
		)

		noti := notification.Notification{
			Type:       "system",
			TargetType: "appointment",
			TargetID:   value.ID,
			IsRead:     false,
			SenderID:   nil,
			ReceiverID: &value.UserID,
			Content:    "Có 1 cuộc hẹn của bạn được hẹn lại vào lúc: " + appointmentDay,
		}

		if err := r.notiService.CreateAndPushSocket(ctx, &noti); err != nil {
			return errors.New("Có lỗi khi thêm thông báo: " + err.Error())
		}
	}

	return nil
}

func (r *AppointmentRepoDB) IsInDay(ctx context.Context, day time.Time) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(&dbmodel.Appointment{}).
		Where("start_time <= ? AND end_time >= ?", day, day).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
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

	for _, value := range dbAppointments {
		appointmentDay := fmt.Sprintf("%02d:%02d %02d/%02d/%04d",
			value.StartTime.Hour(),
			value.StartTime.Minute(),
			value.StartTime.Day(),
			int(value.StartTime.Month()),
			value.StartTime.Year(),
		)

		noti := notification.Notification{
			Type:       "system",
			TargetType: "appointment",
			TargetID:   value.ID,
			IsRead:     false,
			SenderID:   nil,
			ReceiverID: &value.UserID,
			Content:    "Bạn có cuộc hẹn mới vào lúc: " + appointmentDay,
		}

		if err := r.notiService.CreateAndPushSocket(ctx, &noti); err != nil {
			return errors.New("Có lỗi khi thêm thông báo: " + err.Error())
		}
	}

	return nil
}
