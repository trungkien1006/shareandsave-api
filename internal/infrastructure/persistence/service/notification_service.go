package persistence

import (
	"context"
	"final_project/internal/domain/notification"
	"final_project/internal/domain/redis"
	"final_project/internal/domain/user"
	"fmt"
)

type NotificationService struct {
	repo      notification.Repository
	redisRepo redis.Repository
	userRepo  user.Repository
	fcm       notification.Notifier
}

func NewNotificationService(repo notification.Repository, redisRepo redis.Repository, userRepo user.Repository, fcm notification.Notifier) *NotificationService {
	return &NotificationService{
		repo:      repo,
		redisRepo: redisRepo,
		userRepo:  userRepo,
		fcm:       fcm,
	}
}

func (s *NotificationService) CreateAndPushSocket(ctx context.Context, noti *notification.Notification) error {
	var (
		senderName   string = ""
		receiverName string = ""
		err          error
	)

	if err := s.repo.Create(ctx, noti); err != nil {
		return err
	}

	if noti.SenderID != nil {
		senderName, err = s.userRepo.GetUserNameByID(ctx, *noti.SenderID)
		if err != nil {
			return err
		}
	}

	if noti.ReceiverID != nil {
		receiverName, err = s.userRepo.GetUserNameByID(ctx, *noti.ReceiverID)
		if err != nil {
			return err
		}
	}

	notiMap := map[string]interface{}{
		"ID":           noti.ID,
		"senderID":     noti.SenderID,
		"senderName":   senderName,
		"receiverID":   noti.ReceiverID,
		"receiverName": receiverName,
		"type":         noti.Type,
		"targetType":   noti.TargetType,
		"targetID":     noti.TargetID,
		"content":      noti.Content,
		"isRead":       noti.IsRead,
		"createdAt":    noti.CreatedAt,
	}

	fmt.Println("----Thời gian tạo thông báo: " + noti.CreatedAt.String())

	if err := s.redisRepo.InsertToStream(ctx, "notistream", notiMap); err != nil {
		return err
	}

	return nil
}
