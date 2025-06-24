package persistence

import (
	"context"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"time"

	"gorm.io/gorm"
)

type LeaveRequestsRepoDB struct {
	db *gorm.DB
}

func NewLeaveRequestsRepoDB(db *gorm.DB) *LeaveRequestsRepoDB {
	return &LeaveRequestsRepoDB{db: db}
}

func (r *LeaveRequestsRepoDB) IsInLeaveRequest(ctx context.Context, day time.Time) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&dbmodel.LeaveRequests{}).
		Where("start_date <= ? AND end_date >= ?", day, day).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
