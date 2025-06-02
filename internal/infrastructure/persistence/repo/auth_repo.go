package persistence

import "gorm.io/gorm"

type AuthRepoDB struct {
	db *gorm.DB
}

func NewAuthRepoDB(db *gorm.DB) *AuthRepoDB {
	return &AuthRepoDB{db: db}
}
