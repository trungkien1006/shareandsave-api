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
func AppointmentDomainToDB(domain appointment.Appointment, domainItems []appointment.AppointmentItem) Appointment {
	appointmentItems := make([]AppointmentItem, 0)

	for _, value := range domainItems {
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
