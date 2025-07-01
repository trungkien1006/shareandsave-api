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
		"ID":         noti.ID,
		"SenderId":   noti.SenderID,
		"ReceiverID": noti.ReceiverID,
		"Type":       noti.Type,
		"TargetType": noti.TargetType,
		"TargetID":   noti.TargetID,
		"Content":    noti.Content,
		"IsRead":     noti.IsRead,
		"CreatedAt":  noti.CreatedAt,
	}

	if err := s.redisRepo.InsertToStream(ctx, "notistream", notiMap); err != nil {
		return err
	}

	return nil
}
