package dbmodel

import (
	"final_project/internal/domain/notification"
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	SenderID   *uint  `gorm:"index"`
	ReceiverID *uint  `gorm:"index"`
	Title      string `gorm:"size:255"`
	Type       string `gorm:"size:64"`
	TargetType string `gorm:"size:32"`
	TargetID   uint
	Content    string `gorm:"size:255"`
	IsRead     bool
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	Sender   User `gorm:"foreignKey:SenderID"`
	Receiver User `gorm:"foreignKey:ReceiverID"`
}

// Domain to DB
func NotificationDomainToDB(domain notification.Notification) Notification {
	return Notification{
		ID:         domain.ID,
		SenderID:   domain.SenderID,
		ReceiverID: domain.ReceiverID,
		Type:       domain.Type,
		TargetType: domain.TargetType,
		TargetID:   domain.TargetID,
		Content:    domain.Content,
		IsRead:     domain.IsRead,
	}
}

// DB to Domain
func NotificationDBToDomain(db Notification) notification.Notification {
	return notification.Notification{
		ID:           db.ID,
		SenderID:     db.SenderID,
		SenderName:   db.Sender.FullName,
		ReceiverID:   db.ReceiverID,
		ReceiverName: db.Receiver.FullName,
		Type:         db.Type,
		TargetType:   db.TargetType,
		TargetID:     db.TargetID,
		Content:      db.Content,
		IsRead:       db.IsRead,
		CreatedAt:    db.CreatedAt,
	}
}
