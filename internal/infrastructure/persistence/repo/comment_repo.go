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

func (r *CategoryRepoDB) GetAll(ctx context.Context, domainComment *[]comment.Comment, req comment.GetComment) error {
	var (
		query      *gorm.DB
		dbComments []dbmodel.Comment
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Comment{}).
		Where("sender_id = ? AND receiver_id = ?", req.SenderID, req.ReceiverID).
		Where("content LIKE ?", "%"+req.Search+"%").
		Order("created_at DESC")

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
