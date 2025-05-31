package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/post"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type PostRepoDB struct {
	db *gorm.DB
}

func NewPostRepoDB(db *gorm.DB) *PostRepoDB {
	return &PostRepoDB{db: db}
}

func (r *PostRepoDB) Save(ctx context.Context, post *post.Post) error {
	dbPost := dbmodel.PostDomainToDB(*post)

	if err := r.db.Debug().WithContext(ctx).Create(&dbPost).Error; err != nil {
		return errors.New("Lỗi khi tạo bài đăng: " + err.Error())
	}

	*post = dbmodel.PostDBToDomain(dbPost)

	return nil
}
