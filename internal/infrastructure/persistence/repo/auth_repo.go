package persistence

import (
	"context"

	"gorm.io/gorm"
)

type AuthRepoDB struct {
	db *gorm.DB
}

func NewAuthRepoDB(db *gorm.DB) *AuthRepoDB {
	return &AuthRepoDB{db: db}
}

func (r *AuthRepoDB) Login(ctx context.Context, email, password string) error {
	// if err := r.db.Debug().WithContext(ctx).
	// Model(&)

	return nil
}
