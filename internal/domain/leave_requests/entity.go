package leaverequests

import (
	"final_project/internal/pkg/enums"
	"time"
)

type LeaveRequest struct {
	ID        int
	UserID    int
	UserName  string
	LeaveType enums.LeaveRequestType
	Reason    string
	StartDate time.Time
	EndDate   time.Time
	TotalDays float64

	CreatedAt time.Time
}
