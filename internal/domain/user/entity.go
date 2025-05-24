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
	FullName    string `gorm:"size:64"`
	PhoneNumber string `gorm:"unique;size:16"`
	Address     string `gorm:"type:TEXT"`
	Status      int    `gorm:"type:TINYINT"`
	GoodPoint   int    `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (u *User) NewUser(email string, password string, avatar string, fullname string, phoneNumber string, address string, status int, goodPoint int) *User {
	return &User{
		Email:       email,
		Password:    password,
		Avatar:      avatar,
		FullName:    fullname,
		PhoneNumber: phoneNumber,
		Address:     address,
		Status:      status,
		GoodPoint:   goodPoint,
	}
}
