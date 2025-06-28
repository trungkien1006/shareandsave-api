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

func NewNewNotificationService(repo notification.Repository, redisRepo redis.Repository) *NotificationService {
	return &NotificationService{
		repo:      repo,
		redisRepo: redisRepo,
	}
}

func (s *NotificationService) CreateAndPushSocket(ctx context.Context, noti notification.Notification) error {

	return nil
}
