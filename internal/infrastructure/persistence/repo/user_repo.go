package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"fmt"
	"math"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type UserRepoDB struct {
	db *gorm.DB
}

func NewUserRepoDB(db *gorm.DB) *UserRepoDB {
	return &UserRepoDB{db: db}
}

func (r *UserRepoDB) GetAll(ctx context.Context, users *[]user.User, req filter.FilterRequest, roleID uint) (int, error) {
	var (
		query  *gorm.DB
		dbUser []dbmodel.User
	)

	query = r.db.Debug().WithContext(ctx).Where("role_id = ?", roleID).Model(&dbmodel.User{}).Preload("Role")

	if req.SearchBy != "" && req.SearchValue != "" {
		column := strcase.ToSnake(req.SearchBy) // "fullName" -> "full_name"
		query = query.Where(fmt.Sprintf("`%s` LIKE ?", column), "%"+req.SearchValue+"%")
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
		query.Order(strcase.ToSnake(req.Sort) + " " + req.Order)
	}

	if err := query.Find(&dbUser).Error; err != nil {
		return 0, errors.New("Lỗi khi lọc người dùng: " + err.Error())
	}

	//tinh toan total page
	totalPage := int(math.Ceil(float64(totalRecord) / float64(req.Limit)))

	for _, value := range dbUser {
		*users = append(*users, dbmodel.ToDomainUser(value))
	}

	return totalPage, nil
}

func (r *UserRepoDB) IsExist(ctx context.Context, userID uint) (bool, error) {
	var count int64

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.User{}).Where("id = ?", userID).Count(&count).Error; err != nil {
		return false, errors.New("Có lỗi khi kiểm tra user: " + err.Error())
	}

	return count > 0, nil
}

func (r *UserRepoDB) GetClientUserByID(ctx context.Context, domainUser *user.User, userID int, clientID uint) error {
	var dbUser dbmodel.User

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.User{}).Where("id = ?", userID).
		Where("role_id = ?", clientID).
		First(&dbUser).Error; err != nil {
		return errors.New("Lỗi khi tìm kiếm user bằng id: " + err.Error())
	}

	*domainUser = dbmodel.ToDomainUser(dbUser)

	return nil
}

func (r *UserRepoDB) GetAdminUserByID(ctx context.Context, domainUser *user.User, userID int, clientID uint) error {
	var dbUser dbmodel.User

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.User{}).Where("id = ?", userID).
		Where("role_id != ?", clientID).
		Preload("Role").
		Preload("Role.RolePermissions").
		Preload("Role.RolePermissions.Permission").
		First(&dbUser).Error; err != nil {
		return errors.New("Lỗi khi tìm kiếm user bằng id: " + err.Error())
	}

	*domainUser = dbmodel.ToDomainUser(dbUser)

	return nil
}

func (r *UserRepoDB) GetCommonUserByID(ctx context.Context, domainUser *user.User, userID int) error {
	var dbUser dbmodel.User

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.User{}).Where("id = ?", userID).Preload("Role").First(&dbUser).Error; err != nil {
		return errors.New("Lỗi khi tìm kiếm user bằng id: " + err.Error())
	}

	*domainUser = dbmodel.ToDomainUser(dbUser)

	return nil
}

func (r *UserRepoDB) GetByEmailPhoneNumber(ctx context.Context, user *user.User, email string, phoneNumber string) error {
	var dbUser dbmodel.User

	// Truy vấn chỉ lấy trường "id" (không cần toàn bộ user)
	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.User{}).
		Where("email = ? AND phone_number = ?", email, phoneNumber).
		Preload("Role").
		Find(&dbUser).Error; err != nil {
		return errors.New("Lỗi khi tìm kiếm user bằng email và số điện thoại: " + err.Error())
	}

	*user = dbmodel.ToDomainUser(dbUser)

	return nil
}

func (r *UserRepoDB) Save(ctx context.Context, domainUser *user.User) error {
	dbUser := dbmodel.ToDBUser(*domainUser)

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.User{}).Create(&dbUser).Error; err != nil {
		return errors.New("Lỗi khi thêm người dùng mới: " + err.Error())
	}

	*domainUser = dbmodel.ToDomainUser(dbUser)

	return nil
}

func (r *UserRepoDB) Update(ctx context.Context, domainUser *user.User) error {
	dbUser := dbmodel.ToDBUser(*domainUser)

	if err := r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.User{}).
		Omit("CreatedAt").
		Omit("DeleteAt").
		Where("id = ?", domainUser.ID).
		Updates(&dbUser).Error; err != nil {
		return errors.New("Lỗi khi cập nhật người dùng mới: " + err.Error())
	}

	return nil
}

func (r *UserRepoDB) Delete(ctx context.Context, domainUser *user.User) error {
	if err := r.db.Debug().WithContext(ctx).Model(&user.User{}).Delete(&domainUser).Error; err != nil {
		return errors.New("Lỗi khi xóa người dùng: " + err.Error())
	}

	return nil
}

func (r *UserRepoDB) IsEmailExist(ctx context.Context, email string, userID int) (bool, error) {
	var id int64 = 0

	if err := r.db.Debug().WithContext(ctx).Model(&user.User{}).Select("id").Where("email LIKE ?", email).Scan(&id).Error; err != nil {
		return false, errors.New("Lỗi khi kiểm tra email đã tồn tại: " + err.Error())
	}

	if id > 0 {
		if userID == 0 {
			return true, nil
		} else if id == int64(userID) {
			return false, nil
		}
	}

	return false, nil
}

func (r *UserRepoDB) IsPhoneNumberExist(ctx context.Context, phoneNumber string, userID int) (bool, error) {
	var id int64 = 0

	if err := r.db.Debug().WithContext(ctx).Model(&user.User{}).Select("id").Where("phone_number LIKE ?", phoneNumber).Scan(&id).Error; err != nil {
		return false, errors.New("Lỗi khi kiểm tra số điện thoại đã tồn tại: " + err.Error())
	}

	if id > 0 {
		if userID == 0 {
			return true, nil
		} else if id == int64(userID) {
			return false, nil
		}
	}

	return false, nil
}

func (r *UserRepoDB) IsTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&user.User{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
