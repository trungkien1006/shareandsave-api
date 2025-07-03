package persistence

import (
	"context"
	"errors"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type StatisticRepoDB struct {
	db *gorm.DB
}

func NewStatisticRepoDB(db *gorm.DB) *StatisticRepoDB {
	return &StatisticRepoDB{db: db}
}

func (r *StatisticRepoDB) TotalTransaction(ctx context.Context) (int64, int64, error) {
	var (
		total          int64 = 0
		totalLastMonth int64 = 0
	)

	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return 0, 0, errors.New("⚠️ Lỗi khi load location:" + err.Error())
	}

	lastMonth := helpers.GetCurrentTimeVN().AddDate(0, -1, 0)

	start := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, location)
	end := start.AddDate(0, 1, 0) // sang đầu tháng tiếp theo

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Transaction{}).
		Where("status = ?", enums.TransactionStatusSuccess).
		Count(&total).Error; err != nil {
		return total, totalLastMonth, errors.New("Có lỗi khi thống kê tổng số giao dịch: " + err.Error())
	}

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Transaction{}).
		Where("status = ?", enums.TransactionStatusSuccess).
		Where("created_at >= ? AND created_at < ?", start, end).
		Count(&totalLastMonth).Error; err != nil {
		return total, totalLastMonth, errors.New("Có lỗi khi thống kê tổng số giao dịch tháng vừa qua: " + err.Error())
	}

	return total, totalLastMonth, nil
}

func (r *StatisticRepoDB) TotalUser(ctx context.Context, clientID uint) (int64, int64, error) {
	var (
		total          int64 = 0
		totalLastMonth int64 = 0
	)

	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return 0, 0, errors.New("⚠️ Lỗi khi load location:" + err.Error())
	}

	lastMonth := helpers.GetCurrentTimeVN().AddDate(0, -1, 0)

	start := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, location)
	end := start.AddDate(0, 1, 0) // sang đầu tháng tiếp theo

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.User{}).
		Where("status = ? AND role_id = ?", enums.UserStatusActive, clientID).
		Count(&total).Error; err != nil {
		return total, totalLastMonth, errors.New("Có lỗi khi thống kê tổng số thành viên: " + err.Error())
	}

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.User{}).
		Where("status = ? AND role_id = ?", enums.UserStatusActive, clientID).
		Where("created_at >= ? AND created_at < ?", start, end).
		Count(&totalLastMonth).Error; err != nil {
		return total, totalLastMonth, errors.New("Có lỗi khi thống kê tổng số thành viên tháng vừa qua: " + err.Error())
	}

	return total, totalLastMonth, nil
}
