package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/comment"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type CommentRepoDB struct {
	db *gorm.DB
}

func NewCommentRepoDB(db *gorm.DB) *CommentRepoDB {
	return &CommentRepoDB{db: db}
}

func (r *CommentRepoDB) GetAll(ctx context.Context, domainComment *[]comment.Comment, req comment.GetComment) error {
	var (
		query      *gorm.DB
		dbComments []dbmodel.Comment
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Comment{}).
		Where("interest_id = ?", req.InterestID).
		Where("content LIKE ?", "%"+req.Search+"%").
		Order("created_at ASC")

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&dbComments).Error; err != nil {
		return errors.New("Có lỗi khi truy vấn danh sách tin nhắn: " + err.Error())
	}

	for _, value := range dbComments {
		*domainComment = append(*domainComment, dbmodel.CommentDBToDomain(value))
	}

	return nil
}

func (r *CommentRepoDB) Create(ctx context.Context, domainComments *[]comment.Comment) error {
	var dbComments []dbmodel.Comment

	for _, value := range *domainComments {
		dbComments = append(dbComments, dbmodel.CommentDomainToDB(value))
	}

	tx := r.db.Debug().WithContext(ctx)

	// Bulk insert (gorm hỗ trợ)
	err := tx.Model(&dbmodel.Comment{}).Create(&dbComments).Error
	if err != nil {
		return errors.New("Có lỗi khi insert batch comment: " + err.Error())
	}

	return nil
}

func (r *CommentRepoDB) UpdateReadMessage(ctx context.Context, interestID uint) error {
	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Comment{}).
		Where("interest_id = ?", interestID).
		Update("is_read", 1).Error; err != nil {
		return errors.New("Có lỗi kho thực hiện cập nhật trạng thái đã đọc: " + err.Error())
	}

	return nil
}
