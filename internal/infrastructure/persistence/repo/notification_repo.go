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

func (r *NotificationRepoDB) GetAllClient(ctx context.Context, notis *[]notification.Notification, req notification.GetAllNotiRequest, userID uint) (int64, int, error) {
	var (
		query        *gorm.DB
		totalRecords int64
		dbNotis      []dbmodel.Notification
		unreadCount  int64
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Notification{}).
		Where("receiver_id = ?", userID).
		Or("receiver_id IS NULL AND sender_id IS NULL").
		Preload("Sender").
		Preload("Receiver").
		Order("created_at DESC")

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, 0, errors.New("Có lỗi khi đếm tổng số thông báo: " + err.Error())
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&dbNotis).Error; err != nil {
		return 0, 0, errors.New("Có lỗi khi truy suất thông báo: " + err.Error())
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	for _, value := range dbNotis {
		*notis = append(*notis, dbmodel.NotificationDBToDomain(value))
	}

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Notification{}).
		Where("receiver_id = ? AND is_read = 0", userID).
		Or("receiver_id IS NULL AND sender_id IS NULL AND is_read = 0").
		Count(&unreadCount).Error; err != nil {
		return 0, 0, errors.New("Có lỗi khi truy xuất thông báo chưa đọc")
	}

	return unreadCount, totalPages, nil
}

func (r *NotificationRepoDB) GetAllAdmin(ctx context.Context, notis *[]notification.Notification, req notification.GetAllNotiRequest, userID uint) (int64, int, error) {
	var (
		query        *gorm.DB
		totalRecords int64
		dbNotis      []dbmodel.Notification
		unreadCount  int64
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Notification{}).
		Where("receiver_id = ?", userID).
		Preload("Sender").
		Preload("Receiver").
		Order("created_at DESC")

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, 0, errors.New("Có lỗi khi đếm tổng số thông báo: " + err.Error())
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&dbNotis).Error; err != nil {
		return 0, 0, errors.New("Có lỗi khi truy suất thông báo: " + err.Error())
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	for _, value := range dbNotis {
		*notis = append(*notis, dbmodel.NotificationDBToDomain(value))
	}

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Notification{}).
		Where("is_read = 0 AND receiver_id = ?", userID).
		Count(&unreadCount).Error; err != nil {
		return 0, 0, errors.New("Có lỗi khi truy xuất thông báo chưa đọc: " + err.Error())
	}

	return unreadCount, totalPages, nil
}

func (r *NotificationRepoDB) Create(ctx context.Context, noti *notification.Notification) error {
	dbNoti := dbmodel.NotificationDomainToDB(*noti)

	if err := r.db.Debug().WithContext(ctx).
		Create(&dbNoti).Error; err != nil {
		return errors.New("Có lỗi khi tạo thông báo mới: " + err.Error())
	}

	*noti = dbmodel.NotificationDBToDomain(dbNoti)

	return nil
}

func (r *NotificationRepoDB) ReadNoti(ctx context.Context, notiID uint) error {
	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Notification{}).
		Where("id = ?", notiID).
		Update("is_read", 1).Error; err != nil {
		return errors.New("Có lỗi khi cập nhật đã đọc thông báo: " + err.Error())
	}

	return nil
}

func (r *NotificationRepoDB) ReadAllNoti(ctx context.Context, userID uint) error {
	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Notification{}).
		Where("receiver_id = ?", userID).
		Update("is_read", 1).Error; err != nil {
		return errors.New("Có lỗi khi cập nhật đã đọc danh sách thông báo: " + err.Error())
	}

	return nil
}
