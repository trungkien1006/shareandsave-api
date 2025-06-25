package leaverequestapp

import (
	leaverequests "final_project/internal/domain/leave_requests"
)

type UseCase struct {
	repo leaverequests.Repository
}

func NewUseCase(r leaverequests.Repository) *UseCase {
	return &UseCase{repo: r}
}

// func (uc *UseCase) GetAllLeaveRequest(ctx context.Context, leaveReqs *[]leaverequests.LeaveRequest, filter filter.FilterRequest) (int, error) {
// 	totalPage, err := uc.repo.GetAll(ctx, leaveReqs, filter)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return totalPage, nil
// }

// func (uc *UseCase) CreateLeaveRequest(ctx context.Context, domainLeaveRequest leaverequests.LeaveRequest) error {

// 	return nil
// }
