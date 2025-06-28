package persistence

import (
	"gorm.io/gorm"
)

type NotificationRepoDB struct {
	db *gorm.DB
}

func NewNotificationRepoDB(db *gorm.DB) *NotificationRepoDB {
	return &NotificationRepoDB{db: db}
}

// func (r *NotificationRepoDB) Create(ctx context.Context, noti *notification.Notification) error {
// 	var dbNoti dbmodel.Notification

// 	dbNoti = dbmodel.

// 	return nil
// }
