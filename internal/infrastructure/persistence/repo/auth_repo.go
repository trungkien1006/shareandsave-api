package persistence

import (
	"context"
	"errors"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type AuthRepoDB struct {
	db *gorm.DB
}

func NewAuthRepoDB(db *gorm.DB) *AuthRepoDB {
	return &AuthRepoDB{db: db}
}

func (r *AuthRepoDB) Login(ctx context.Context, email, hashedPassword string) error {
	var dbUser dbmodel.User

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.User{}).
		Where("email = ?", hashedPassword).
		First(&dbUser).Error; err != nil {
		return errors.New("Email không tồn tại: " + err.Error())
	}

	return nil
}
