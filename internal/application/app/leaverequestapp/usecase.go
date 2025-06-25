package leaverequestapp

import (
	"context"
	"final_project/internal/domain/filter"
	leaverequests "final_project/internal/domain/leave_requests"
)

type UseCase struct {
	repo leaverequests.Repository
}

func NewUseCase(r leaverequests.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllLeaveRequest(ctx context.Context, leaveReq leaverequests.LeaveRequest, filter filter.FilterRequest) (int, error) {
	return 0, nil
}
