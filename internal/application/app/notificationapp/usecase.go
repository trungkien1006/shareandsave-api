package notificationapp

import (
	"context"
	"final_project/internal/domain/notification"
	"final_project/internal/domain/redis"
	"time"

	redisV9 "github.com/redis/go-redis/v9"
)

type UseCase struct {
	repo      notification.Repository
	service   notification.Service
	redisRepo redis.Repository
}

func NewUseCase(r notification.Repository, service notification.Service, redisRepo redis.Repository) *UseCase {
	return &UseCase{repo: r, service: service, redisRepo: redisRepo}
}

func (uc *UseCase) StoreFCMToken(ctx context.Context, token, userID string) error {
	redisToken, err := uc.redisRepo.GetFromRedis(ctx, "user:"+userID+":fcmToken")
	if err != nil {
		if err == redisV9.Nil {
			redisToken = ""
		} else {
			return err
		}
	}

	if redisToken != token {
		if err := uc.redisRepo.InsertToRedis(ctx, "user:"+userID+":fcmToken", token, 24*30*time.Hour); err != nil {
			return err
		}

		if err := uc.redisRepo.InsertToRedis(ctx, "fcmToken:"+token, userID, 24*30*time.Hour); err != nil {
			return err
		}
	} else {
		if err := uc.redisRepo.InsertToRedis(ctx, "fcmToken:"+token, userID, 24*30*time.Hour); err != nil {
			return err
		}
	}

	return nil
}

func (uc *UseCase) DeleteFCMToken(ctx context.Context, token, userID string) error {
	if err := uc.redisRepo.DeleteFromRedis(ctx, "user:"+userID+":fcmToken"); err != nil {
		return err
	}

	if err := uc.redisRepo.DeleteFromRedis(ctx, "fcmToken:"+token); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetAllClientNoti(ctx context.Context, notis *[]notification.Notification, req notification.GetAllNotiRequest, userID uint) (int64, int, error) {
	unreadCount, totalPage, err := uc.repo.GetAllClient(ctx, notis, req, userID)
	if err != nil {
		return 0, 0, err
	}

	return unreadCount, totalPage, nil
}

func (uc *UseCase) GetAllAdminNoti(ctx context.Context, notis *[]notification.Notification, req notification.GetAllNotiRequest, userID uint) (int64, int, error) {
	unreadCount, totalPage, err := uc.repo.GetAllAdmin(ctx, notis, req, userID)
	if err != nil {
		return 0, 0, err
	}

	return unreadCount, totalPage, nil
}

func (uc *UseCase) CreateSystemNoti(ctx context.Context, noti *notification.Notification) error {
	if err := uc.service.CreateAndPushSocket(ctx, noti); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) ReadNoti(ctx context.Context, notiID uint) error {
	if err := uc.repo.ReadNoti(ctx, notiID); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) ReadAllNoti(ctx context.Context, userID uint) error {
	if err := uc.repo.ReadAllNoti(ctx, userID); err != nil {
		return err
	}

	return nil
}
