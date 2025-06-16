package persistence

import "gorm.io/gorm"

type CommentRepoDB struct {
	db *gorm.DB
}

func NewCommentRepoDB(db *gorm.DB) *CommentRepoDB {
	return &CommentRepoDB{db: db}
}
