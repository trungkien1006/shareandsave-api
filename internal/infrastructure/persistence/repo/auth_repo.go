package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/user"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"

	"gorm.io/gorm"
)

type AuthRepoDB struct {
	db *gorm.DB
}

func NewAuthRepoDB(db *gorm.DB) *AuthRepoDB {
	return &AuthRepoDB{db: db}
}

func (r *AuthRepoDB) Login(ctx context.Context, user *user.User, email, password string) ([]string, error) {
	var (
		dbUser      dbmodel.User
		permisisons []string
	)

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.User{}).
		Where("email = ?", email).
		Preload("Role").First(&dbUser).Error; err != nil {
		return nil, errors.New("Email không tồn tại: " + err.Error())
	}

	if !hash.CheckPasswordHash(password, dbUser.Password) {
		return nil, errors.New("Mật khẩu không đúng")
	}

	if dbUser.Status == int8(enums.UserStatusLocked) {
		return nil, errors.New("Tài khoản đã bị khóa")
	}

	if dbUser.Role.Name != "Client" {
		if err := r.db.Debug().WithContext(ctx).
			Model(&dbmodel.Permission{}).
			Table("permission as permission").
			Select("permission.code").
			Joins("join role_permission on role_permission.permission_id = permission.id").
			Where("role_permission.role_id = ?", dbUser.Role.ID).
			Scan(&permisisons).Error; err != nil {
			return nil, errors.New("Lỗi khi truy suất danh sách quyền: " + err.Error())
		}
	}

	*user = dbmodel.ToDomainUser(dbUser)

	return permisisons, nil
}
