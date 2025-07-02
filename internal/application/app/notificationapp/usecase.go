package notificationapp

import (
	"context"
	"final_project/internal/domain/notification"
)

type UseCase struct {
	repo notification.Repository
}

func NewUseCase(r notification.Repository) *UseCase {
	return &UseCase{repo: r}
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
