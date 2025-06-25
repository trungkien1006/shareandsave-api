package leaverequests

type Repository interface {
	// GetAll(ctx context.Context, leaveRequests *[]LeaveRequest, req filter.FilterRequest) (int, error)
	// GetByID(ctx context.Context, leaveRequest *LeaveRequest, leaveReqID uint) error
	// Create(ctx context.Context, leaveReq LeaveRequest) error
	// IsInLeaveRequest(ctx context.Context, day time.Time) (bool, error)
}
