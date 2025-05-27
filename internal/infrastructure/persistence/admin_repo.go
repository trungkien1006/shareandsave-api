package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/admin"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/role_permission"
	"final_project/internal/pkg/helpers"
	"final_project/internal/reference"
	"math"

	"gorm.io/gorm"
)

type AdminRepoDB struct {
	db *gorm.DB
}

func NewAdminRepoDB(db *gorm.DB) *AdminRepoDB {
	return &AdminRepoDB{db: db}
}

func (r *AdminRepoDB) GetAll(ctx context.Context, admins *[]admin.Admin, req filter.FilterRequest) (int, error) {
	var tableName = "admin"
	var query *gorm.DB

	query = r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Preload("Role")

	if req.Filter != "" {
		var filters []reference.FilterStruc

		err := json.Unmarshal([]byte(req.Filter), &filters)
		if err != nil {
			return 0, errors.New("Lỗi khi chuyển đổi filter từ JSON thành struct: " + err.Error())
		}

		helpers.Filter(query, filters, tableName)
	}

	var totalRecord int64 = 0

	if err := query.Count(&totalRecord).Error; err != nil {
		return 0, errors.New("Lỗi khi đếm tổng số record của admin: " + err.Error())
	}

	if req.Limit != 0 && req.Page != 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if req.Sort != "" {
		query.Order(req.Sort + " " + req.Order)
	}

	if err := query.Find(&admins).Error; err != nil {
		return 0, errors.New("Lỗi khi lọc admin: " + err.Error())
	}

	totalPage := int(math.Ceil(float64(totalRecord) / float64(req.Limit)))
	return totalPage, nil
}

func (r *AdminRepoDB) GetByID(ctx context.Context, domainAdmin *admin.Admin, adminID int) error {
	if err := r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Preload("Role").Where("id = ?", adminID).First(&domainAdmin).Error; err != nil {
		return errors.New("Lỗi khi tìm kiếm admin bằng id: " + err.Error())
	}
	return nil
}

func (r *AdminRepoDB) Save(ctx context.Context, domainAdmin *admin.Admin) error {
	if err := r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Create(&domainAdmin).Error; err != nil {
		return errors.New("Lỗi khi thêm admin mới: " + err.Error())
	}
	return nil
}

func (r *AdminRepoDB) Update(ctx context.Context, domainAdmin *admin.Admin) error {
	if err := r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Save(&domainAdmin).Error; err != nil {
		return errors.New("Lỗi khi cập nhật admin: " + err.Error())
	}
	return nil
}

func (r *AdminRepoDB) Delete(ctx context.Context, domainAdmin *admin.Admin) error {
	if err := r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Delete(&domainAdmin).Error; err != nil {
		return errors.New("Lỗi khi xóa admin: " + err.Error())
	}
	return nil
}

func (r *AdminRepoDB) IsEmailExist(ctx context.Context, email string) (bool, error) {
	var count int64 = 0
	if err := r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Where("email LIKE ?", email).Count(&count).Error; err != nil {
		return false, errors.New("Lỗi khi kiểm tra email đã tồn tại: " + err.Error())
	}
	return count > 0, nil
}

func (r *AdminRepoDB) IsRoleExist(ctx context.Context, roleId uint) (bool, error) {
	var count int64 = 0
	if err := r.db.Debug().WithContext(ctx).Model(&role_permission.Role{}).Where("id = ?", roleId).Count(&count).Error; err != nil {
		return false, errors.New("Lỗi khi kiểm tra chức vụ đã tồn tại: " + err.Error())
	}
	return count > 0, nil
}

func (r *AdminRepoDB) IsTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&admin.Admin{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
