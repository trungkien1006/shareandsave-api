package persistence

import (
	"context"
	"final_project/internal/domain/request"

	"gorm.io/gorm"
)

type RequestRepoDB struct {
	db *gorm.DB
}

func NewRequestRepoDB(db *gorm.DB) *RequestRepoDB {
	return &RequestRepoDB{db: db}
}

func (r *RequestRepoDB) Create(ctx context.Context, request *request.SendRequest) error {
	if err := r.db.Create(&request).Error; err != nil {
		return err
	}

	return nil
}
