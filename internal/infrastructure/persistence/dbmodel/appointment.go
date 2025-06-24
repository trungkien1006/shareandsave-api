package dbmodel

import (
	"final_project/internal/domain/appointment"
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint `gorm:"index"`
	StartTime time.Time
	EndTime   time.Time
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:UserID"`

	AppointmentItem []AppointmentItem `gorm:"foreignKey:AppointmentID"`
}

// Domain to DB
func AppointmentDomainToDB(domain appointment.Appointment) Appointment {
	appointmentItems := make([]AppointmentItem, 0)

	for _, value := range domain.AppointmentItems {
		appointmentItems = append(appointmentItems, AppointmentItem{
			ItemID:          value.ItemID,
			ActualQuantity:  value.ActualQuantity,
			MissingQuantity: value.MissingQuantity,
		})
	}

	return Appointment{
		UserID:          domain.UserID,
		StartTime:       domain.StartTime,
		EndTime:         domain.EndTime,
		Status:          domain.Status,
		AppointmentItem: appointmentItems,
	}
}

// DB to Domain
func AppointmentDBToDomain(db Appointment) appointment.Appointment {
	var items []appointment.AppointmentItem

	for _, value := range db.AppointmentItem {
		items = append(items, appointment.AppointmentItem{
			ID:              value.ID,
			AppointmentID:   value.AppointmentID,
			ItemID:          value.ItemID,
			ActualQuantity:  value.ActualQuantity,
			MissingQuantity: value.MissingQuantity,
		})
	}

	return appointment.Appointment{
		ID:               db.ID,
		UserID:           db.UserID,
		StartTime:        db.StartTime,
		EndTime:          db.EndTime,
		Status:           db.Status,
		AppointmentItems: items,
		CreatedAt:        db.CreatedAt,
	}
}
