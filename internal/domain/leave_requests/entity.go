package leaverequests

import "time"

type LeaveRequest struct {
	ID        int
	UserID    int
	UserName  string
	LeaveType int
	Reason    string
	StartDate time.Time
	EndDate   time.Time
	TotalDays float64

	CreatedAt time.Time
}
