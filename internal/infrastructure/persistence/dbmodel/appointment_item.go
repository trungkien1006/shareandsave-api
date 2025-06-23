package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type AppointmentItem struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	AppointmentID   uint `gorm:"index"`
	ItemID          uint `gorm:"index"`
	ActualQuantity  int
	MissingQuantity int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	// Quan hệ
	Appointment Appointment `gorm:"foreignKey:AppointmentID"`
	Item        Item        `gorm:"foreignKey:ItemID"`
}
