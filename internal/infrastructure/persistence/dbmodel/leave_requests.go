package dbmodel

import (
	leaverequests "final_project/internal/domain/leave_requests"
	"final_project/internal/pkg/enums"
	"time"

	"gorm.io/gorm"
)

type LeaveRequests struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	LeaveType int       `gorm:"type:INT;not null" json:"leave_type"`
	Reason    string    `gorm:"type:text;not null" json:"reason"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time `gorm:"type:date;not null" json:"end_date"`
	TotalDays float64   `gorm:"not null" json:"total_days"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Khóa ngoại - quan hệ 1-1 tới User
	User User `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
}

// DB to Domain
func LeaveReqDBToDomain(db LeaveRequests) leaverequests.LeaveRequest {
	return leaverequests.LeaveRequest{
		ID:        db.ID,
		UserID:    db.UserID,
		UserName:  db.User.FullName,
		LeaveType: enums.LeaveRequestType(db.LeaveType),
		Reason:    db.Reason,
		StartDate: db.StartDate,
		EndDate:   db.EndDate,
		TotalDays: db.TotalDays,
		CreatedAt: db.CreatedAt,
	}
}

// Domain to DB
func LeaveReqDomainToDB(domain leaverequests.LeaveRequest) LeaveRequests {
	return LeaveRequests{
		ID:        domain.ID,
		UserID:    domain.UserID,
		LeaveType: int(domain.LeaveType),
		Reason:    domain.Reason,
		StartDate: domain.StartDate,
		EndDate:   domain.EndDate,
		TotalDays: domain.TotalDays,
		CreatedAt: domain.CreatedAt,
	}
}
