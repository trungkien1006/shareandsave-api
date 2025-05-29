package persistence

import (
	"context"
	"final_project/internal/domain/request"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type RequestRepoDB struct {
	db *gorm.DB
}

func NewRequestRepoDB(db *gorm.DB) *RequestRepoDB {
	return &RequestRepoDB{db: db}
}

func (r *RequestRepoDB) Create(ctx context.Context, request *request.SendRequest) error {
	var dbRequest = dbmodel.RequestDomainToDB(*request)

	if err := r.db.Create(&dbRequest).Error; err != nil {
		return err
	}

	*request = dbmodel.SendRequestToDomain(dbRequest)

	return nil
}
