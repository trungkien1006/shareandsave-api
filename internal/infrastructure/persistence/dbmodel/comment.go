package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	InterestID uint   `gorm:"index"`
	SenderID   uint   `gorm:"index"`
	ReceiverID uint   `gorm:"index"`
	Content    string `gorm:"type:TEXT"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	Interest Interest `gorm:"foreignKey:InterestID"`
	Sender   User     `gorm:"foreignKey:SenderID"`
	Receiver User     `gorm:"foreignKey:ReceiverID"`
}
