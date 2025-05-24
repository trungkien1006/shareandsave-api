package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/helpers"
	"final_project/internal/reference"
	"math"

	"gorm.io/gorm"
)

type UserRepoDB struct {
	db *gorm.DB
}

func NewUserRepoDB(db *gorm.DB) *UserRepoDB {
	return &UserRepoDB{db: db}
}

func (r *UserRepoDB) GetAll(ctx context.Context, users *[]user.User, req filter.FilterRequest) (int, error) {
	var tableName = "user"
	var query *gorm.DB

	query = r.db.Debug().WithContext(ctx).Table(tableName)

	if req.Filter != "" {
		var filters []reference.FilterStruc

		err := json.Unmarshal([]byte(req.Filter), &filters)
		if err != nil {
			return 0, errors.New("Lỗi khi chuyển đổi filter từ JSON thành struct: " + err.Error())
		}

		helpers.Filter(query, filters, tableName)
	}

	var totalRecord int64 = 0

	//lay ra tong so record
	if err := query.Count(&totalRecord).Error; err != nil {
		return 0, errors.New("Lỗi khi đếm tổng số record của user: " + err.Error())
	}

	//phan trang
	if req.Limit != 0 && req.Page != 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	//sort du lieu
	if req.Sort != "" {
		query.Order(req.Sort + " " + req.Order)
	}

	//tinh toan total page
	totalPage := int(math.Ceil(float64(totalRecord) / float64(req.Limit)))

	return totalPage, nil
}

func (r *UserRepoDB) GetByID(ctx context.Context, user *user.User, user_id int) error {
	if err := r.db.Debug().WithContext(ctx).Where("id = ?", user_id).First(&user).Error; err != nil {
		return errors.New("Lỗi khi tìm kiếm user bằng id: " + err.Error())
	}

	return nil
}

func (r *UserRepoDB) Save(ctx context.Context, user *user.User) error {
	if err := r.db.Debug().WithContext(ctx).Create(&user).Error; err != nil {
		return errors.New("Lỗi khi thêm người dùng mới: " + err.Error())
	}

	return nil
}

func (r *UserRepoDB) Update(ctx context.Context, user *user.User) error {
	if err := r.db.Debug().WithContext(ctx).Save(&user).Error; err != nil {
		return errors.New("Lỗi khi cập nhật người dùng mới: " + err.Error())
	}

	return nil
}

func (r *UserRepoDB) IsEmailExist(ctx context.Context, email string) (bool, error) {
	var count int64 = 0

	if err := r.db.Debug().WithContext(ctx).Where("email LIKE ?", email).Count(&count).Error; err != nil {
		return false, errors.New("Lỗi khi kiểm tra email đã tồn tại: " + err.Error())
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *UserRepoDB) IsPhoneNumberExist(ctx context.Context, phoneNumber string) (bool, error) {
	var count int64 = 0

	if err := r.db.Debug().WithContext(ctx).Where("phone_number LIKE ?", phoneNumber).Count(&count).Error; err != nil {
		return false, errors.New("Lỗi khi kiểm tra số điện thoại đã tồn tại: " + err.Error())
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
