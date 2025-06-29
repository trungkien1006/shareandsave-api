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

func (r *NotificationRepoDB) GetAll(ctx context.Context, notis *[]notification.Notification, req notification.GetAllNotiRequest, userID uint) (int, error) {
	var (
		query        *gorm.DB
		totalRecords int64
		dbNotis      []dbmodel.Notification
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Notification{}).
		Where("receiver_id = ?", userID).
		Order("created_at DESC")

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, err
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&dbNotis).Error; err != nil {
		return 0, err
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	for _, value := range dbNotis {
		*notis = append(*notis, dbmodel.NotificationDBToDomain(value))
	}

	return totalPages, nil
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
