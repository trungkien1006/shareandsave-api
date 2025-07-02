package notificationapp

import (
	"context"
	"final_project/internal/domain/notification"
)

type UseCase struct {
	repo    notification.Repository
	service notification.Service
}

func NewUseCase(r notification.Repository, service notification.Service) *UseCase {
	return &UseCase{repo: r, service: service}
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
