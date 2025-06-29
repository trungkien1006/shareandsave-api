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

func (uc *UseCase) GetAllNoti(ctx context.Context, notis *[]notification.Notification, req notification.GetAllNotiRequest, userID uint) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, notis, req, userID)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}
