package persistence

import (
	"context"
	"final_project/internal/domain/notification"
	"final_project/internal/domain/redis"
)

type NotificationService struct {
	repo      notification.Repository
	redisRepo redis.Repository
}

func NewNotificationService(repo notification.Repository, redisRepo redis.Repository) *NotificationService {
	return &NotificationService{
		repo:      repo,
		redisRepo: redisRepo,
	}
}

func (s *NotificationService) CreateAndPushSocket(ctx context.Context, noti *notification.Notification) error {
	if err := s.repo.Create(ctx, noti); err != nil {
		return err
	}

	notiMap := map[string]interface{}{
		"iD":         noti.ID,
		"senderId":   noti.SenderID,
		"receiverID": noti.ReceiverID,
		"type":       noti.Type,
		"targetType": noti.TargetType,
		"targetID":   noti.TargetID,
		"content":    noti.Content,
		"isRead":     noti.IsRead,
		"createdAt":  noti.CreatedAt,
	}

	if err := s.redisRepo.InsertToStream(ctx, "notistream", notiMap); err != nil {
		return err
	}

	return nil
}
