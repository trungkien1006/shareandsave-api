package notificationapp

import "final_project/internal/domain/notification"

type UseCase struct {
	repo notification.Repository
}

func NewUseCase(r notification.Repository) *UseCase {
	return &UseCase{repo: r}
}
