package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/admin"
	"final_project/internal/domain/filter"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"math"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type AdminRepoDB struct {
	db *gorm.DB
}

func NewAdminRepoDB(db *gorm.DB) *AdminRepoDB {
	return &AdminRepoDB{db: db}
}

func (r *AdminRepoDB) GetAllWithRole(ctx context.Context, filter filter.FilterRequest) ([]admin.AdminWithRole, int, error) {
	var (
		query  *gorm.DB
		admins []dbmodel.Admin
	)

	query = r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Preload("Role")

	if filter.SearchBy != "" && filter.SearchValue != "" {
		query = query.Where(strcase.ToSnake(filter.SearchBy)+" LIKE ?", "%"+filter.SearchValue+"%")
	}

	var totalRecord int64 = 0

	if err := query.Count(&totalRecord).Error; err != nil {
		return []admin.AdminWithRole{}, 0, errors.New("Lỗi khi đếm tổng số record của admin: " + err.Error())
	}

	if filter.Limit != 0 && filter.Page != 0 {
		query.Offset((filter.Page - 1) * filter.Limit).Limit(filter.Limit)
	}

	if filter.Sort != "" {
		query.Order(strcase.ToSnake(filter.Sort) + " " + filter.Order)
	}

	if err := query.Find(&admins).Preload("Role").Error; err != nil {
		return []admin.AdminWithRole{}, 0, errors.New("Lỗi khi lọc admin: " + err.Error())
	}

	totalPage := int(math.Ceil(float64(totalRecord) / float64(filter.Limit)))

	var result []admin.AdminWithRole
	for _, dbAdmin := range admins {
		result = append(result, admin.AdminWithRole{
			Admin: admin.Admin{
				ID:       dbAdmin.ID,
				Email:    dbAdmin.Email,
				FullName: dbAdmin.FullName,
				Status:   dbAdmin.Status,
				RoleID:   dbAdmin.RoleID,
			},
			RoleName: dbAdmin.Role.Name,
		})
	}
	return result, totalPage, nil
}

func (r *AdminRepoDB) GetAll(ctx context.Context, admins *[]admin.Admin, req filter.FilterRequest) (int, error) {
	var query *gorm.DB

	query = r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Preload("Role")

	if req.SearchBy != "" && req.SearchValue != "" {
		query = query.Where(strcase.ToSnake(req.SearchBy)+" LIKE ?", "%"+req.SearchValue+"%")
	}

	var totalRecord int64 = 0

	if err := query.Count(&totalRecord).Error; err != nil {
		return 0, errors.New("Lỗi khi đếm tổng số record của admin: " + err.Error())
	}

	if req.Limit != 0 && req.Page != 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if req.Sort != "" {
		query.Order(strcase.ToSnake(req.Sort) + " " + req.Order)
	}

	if err := query.Find(&admins).Error; err != nil {
		return 0, errors.New("Lỗi khi lọc admin: " + err.Error())
	}

	totalPage := int(math.Ceil(float64(totalRecord) / float64(req.Limit)))
	return totalPage, nil
}

func (r *AdminRepoDB) GetByID(ctx context.Context, domainAdmin *admin.Admin, adminID int) error {
	var dbAdmin dbmodel.Admin

	if err := r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Preload("Role").Where("id = ?", adminID).First(&dbAdmin).Error; err != nil {
		return errors.New("Lỗi khi tìm kiếm admin bằng id: " + err.Error())
	}

	*domainAdmin = dbmodel.AdminDBToDomain(dbAdmin)

	return nil
}

func (r *AdminRepoDB) GetByIDWithRole(ctx context.Context, adminID uint) (admin.AdminWithRole, error) {
	var dbAdmin dbmodel.Admin

	if err := r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Preload("Role").Where("id = ?", adminID).First(&dbAdmin).Error; err != nil {
		return admin.AdminWithRole{}, errors.New("Lỗi khi tìm kiếm admin bằng id: " + err.Error())
	}

	result := admin.AdminWithRole{
		Admin: admin.Admin{ID: dbAdmin.ID,
			Email:    dbAdmin.Email,
			FullName: dbAdmin.FullName,
			Status:   dbAdmin.Status,
			RoleID:   dbAdmin.RoleID,
		},
		RoleName: dbAdmin.Role.Name,
	}

	return result, nil
}

func (r *AdminRepoDB) Save(ctx context.Context, domainAdmin admin.Admin) (admin.Admin, error) {
	dbAdmin := dbmodel.AdminDomainToDB(domainAdmin)

	if err := r.db.Debug().WithContext(ctx).Create(&dbAdmin).Error; err != nil {
		return admin.Admin{}, errors.New("Lỗi khi lưu admin: " + err.Error())
	}

	return dbmodel.AdminDBToDomain(dbAdmin), nil
}

func (r *AdminRepoDB) Update(ctx context.Context, domainAdmin *admin.Admin) error {
	var dbAdmin dbmodel.Admin

	dbAdmin = dbmodel.AdminDomainToDB(*domainAdmin)

	if err := r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Where("id = ?", domainAdmin.ID).Save(&dbAdmin).Error; err != nil {
		return errors.New("Lỗi khi cập nhật admin: " + err.Error())
	}

	*domainAdmin = dbmodel.AdminDBToDomain(dbAdmin)

	return nil
}

func (r *AdminRepoDB) Delete(ctx context.Context, domainAdmin *admin.Admin) error {
	var dbAdmin dbmodel.Admin

	dbAdmin = dbmodel.AdminDomainToDB(*domainAdmin)

	if err := r.db.Debug().WithContext(ctx).Model(&admin.Admin{}).Delete(&dbAdmin).Error; err != nil {
		return errors.New("Lỗi khi xóa admin: " + err.Error())
	}

	*domainAdmin = dbmodel.AdminDBToDomain(dbAdmin)

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
	if err := r.db.Debug().WithContext(ctx).Model(&rolepermission.Role{}).Where("id = ?", roleId).Count(&count).Error; err != nil {
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
