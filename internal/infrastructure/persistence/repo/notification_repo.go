package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/notification"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type NotificationRepoDB struct {
	db *gorm.DB
}

func NewNotificationRepoDB(db *gorm.DB) *NotificationRepoDB {
	return &NotificationRepoDB{db: db}
}

func (r *NotificationRepoDB) Create(ctx context.Context, noti *notification.Notification) error {
	dbNoti := dbmodel.NotificationDomainToDB(*noti)

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Notification{}).
		Create(&dbNoti).Error; err != nil {
		return errors.New("Có lỗi khi tạo thông báo mới: " + err.Error())
	}

	*noti = dbmodel.NotificationDBToDomain(dbNoti)

	return nil
}
