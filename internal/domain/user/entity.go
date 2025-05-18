package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Email       string `gorm:"unique;size:255;not null"`
	Password    string `gorm:"size:255;not null"`
	Avatar      string `gorm:"type:LONGTEXT"`
	Active      bool
	Fullname    string `gorm:"size:64"`
	PhoneNumber string `gorm:"unique;size:16"`
	Address     string `gorm:"type:TEXT"`
	Status      int8   `gorm:"type:TINYINT"`
	GoodPoint   int    `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func NewUser(email string, password string, avatar string, fullname string, phoneNumber string, address string, active bool, status int8) *User {
	return &User{
		Email:       email,
		Password:    password,
		Avatar:      avatar,
		Fullname:    fullname,
		PhoneNumber: phoneNumber,
		Address:     address,
		Active:      active,
		Status:      status,
	}
}
