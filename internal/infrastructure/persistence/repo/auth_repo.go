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

func (r *AuthRepoDB) Login(ctx context.Context, user *user.User, email, password string) error {
	var dbUser dbmodel.User

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.User{}).
		Where("email = ?", email).
		First(&dbUser).Error; err != nil {
		return errors.New("Email không tồn tại: " + err.Error())
	}

	if !hash.CheckPasswordHash(password, dbUser.Password) {
		return errors.New("Mật khẩu không đúng")
	}

	if dbUser.Status == int8(enums.UserStatusLocked) {
		return errors.New("Tài khoản đã bị khóa")
	}

	*user = dbmodel.ToDomainUser(dbUser)

	return nil
}
