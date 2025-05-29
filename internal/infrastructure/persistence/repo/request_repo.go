package persistence

import (
	"context"
	sendrequest "final_project/internal/domain/send_request"

	"gorm.io/gorm"
)

type SendRequestRepoDB struct {
	db *gorm.DB
}

func NewSendRequestRepoDB(db *gorm.DB) *SendRequestRepoDB {
	return &SendRequestRepoDB{db: db}
}

func (r *SendRequestRepoDB) Create(ctx context.Context, request *sendrequest.SendRequest) error {
	if err := r.db.Create(&request).Error; err != nil {
		return err
	}

	return nil
}
