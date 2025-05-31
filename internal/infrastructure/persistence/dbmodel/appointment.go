package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	UserID        uint `gorm:"index"`
	ScheduledTime time.Time
	Status        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:UserID"`

	AppointmentItemWarehouses []AppointmentItemWarehouse `gorm:"foreignKey:AppointmentID"`
}
