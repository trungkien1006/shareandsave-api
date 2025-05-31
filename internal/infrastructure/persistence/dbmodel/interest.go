package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Interest struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint `gorm:"index"`
	PostID    uint `gorm:"index"`
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:UserID"`
	Post Post `gorm:"foreignKey:PostID"`

	Comments     []Comment     `gorm:"foreignKey:InterestID"`
	Transactions []Transaction `gorm:"foreignKey:InterestID"`
}
