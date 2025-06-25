package handler

import "final_project/internal/application/app/leaverequestapp"

type LeaveRequestsHandler struct {
	uc *leaverequestapp.UseCase
}

func NewLeaveRequestsHandler(uc *leaverequestapp.UseCase) *LeaveRequestsHandler {
	return &LeaveRequestsHandler{uc: uc}
}
