package persistence

import (
	"context"
	"errors"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type StatisticRepoDB struct {
	db *gorm.DB
}

func NewStatisticRepoDB(db *gorm.DB) *StatisticRepoDB {
	return &StatisticRepoDB{db: db}
}

func (r *StatisticRepoDB) StatisticTransactionInYear(ctx context.Context, year string) ([]int64, error) {
	totals := make([]int64, 12) // Mỗi phần tử tương ứng 1 tháng

	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return totals, errors.New("⚠️ Lỗi khi load location: " + err.Error())
	}

	// Parse chuỗi năm thành số nguyên
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return totals, errors.New("⚠️ Năm không hợp lệ: " + err.Error())
	}

	// Lặp qua 12 tháng
	for month := 1; month <= 12; month++ {
		start := time.Date(yearInt, time.Month(month), 1, 0, 0, 0, 0, location)
		end := start.AddDate(0, 1, 0) // đầu tháng tiếp theo

		var count int64
		if err := r.db.WithContext(ctx).
			Model(&dbmodel.Transaction{}).
			Where("status = ?", enums.TransactionStatusSuccess).
			Where("created_at >= ? AND created_at < ?", start, end).
			Count(&count).Error; err != nil {
			return totals, errors.New("Có lỗi khi thống kê giao dịch tháng " + strconv.Itoa(month) + "/" + strconv.Itoa(yearInt) + ": " + err.Error())
		}

		totals[month-1] = count
	}

	return totals, nil
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

func (r *StatisticRepoDB) TotalPost(ctx context.Context) (int64, int64, error) {
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
		Model(&dbmodel.Post{}).
		Where("status = ? OR status = ?", enums.PostStatusApproved, enums.PostStatusSeal).
		Count(&total).Error; err != nil {
		return total, totalLastMonth, errors.New("Có lỗi khi thống kê tổng số bài viết: " + err.Error())
	}

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Post{}).
		Where("(status = ? OR status = ?) AND (created_at >= ? AND created_at < ?)", enums.PostStatusApproved, enums.PostStatusSeal, start, end).
		Count(&totalLastMonth).Error; err != nil {
		return total, totalLastMonth, errors.New("Có lỗi khi thống kê tổng số bài viết tháng vừa qua: " + err.Error())
	}

	return total, totalLastMonth, nil
}

func (r *StatisticRepoDB) TotalItemWarehouse(ctx context.Context) (int64, int64, error) {
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
		Model(&dbmodel.ItemWarehouse{}).
		Where("status = ?", enums.ItemWarehouseStatusInStock).
		Count(&total).Error; err != nil {
		return total, totalLastMonth, errors.New("Có lỗi khi thống kê tổng số hàng tồn: " + err.Error())
	}

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.ItemWarehouse{}).
		Where("status = ?", enums.ItemWarehouseStatusInStock).
		Where("created_at >= ? AND created_at < ?", start, end).
		Count(&totalLastMonth).Error; err != nil {
		return total, totalLastMonth, errors.New("Có lỗi khi thống kê tổng số hàng tồn tháng vừa qua: " + err.Error())
	}

	return total, totalLastMonth, nil
}
