package leaverequestapp

import leaverequests "final_project/internal/domain/leave_requests"

type UseCase struct {
	repo leaverequests.Repository
}

func NewUseCase(r leaverequests.Repository) *UseCase {
	return &UseCase{repo: r}
}
