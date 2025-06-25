package leaverequestdto

import (
	"time"
)

type LeaveRequestDTO struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	UserName  string    `json:"userName"`
	LeaveType int       `json:"leaveType"`
	Reason    string    `json:"reason"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	TotalDays float64   `json:"totalDays"`

	CreatedAt time.Time `json:"createdAt"`
}
