package notification

import (
	"os/user"
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	SenderID   uint   `gorm:"index"`
	ReceiverID uint   `gorm:"index"`
	Type       string `gorm:"size:64"`
	TargetType string `gorm:"size:32"`
	TargetID   uint
	Message    string `gorm:"size:255"`
	IsRead     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Sender     user.User      `gorm:"foreignKey:SenderID"`
	Receiver   user.User      `gorm:"foreignKey:ReceiverID"`
}

func NewNotification(senderID, receiverID uint, nType, targetType string, targetID uint, message string, isRead bool) *Notification {
	return &Notification{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Type:       nType,
		TargetType: targetType,
		TargetID:   targetID,
		Message:    message,
		IsRead:     isRead,
	}
}
