package dbmodel

import (
	"final_project/internal/domain/comment"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	InterestID uint   `gorm:"index"`
	SenderID   uint   `gorm:"index"`
	ReceiverID uint   `gorm:"index"`
	Content    string `gorm:"type:TEXT"`
	IsRead     uint   `gorm:"type:TINYINT"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	Interest Interest `gorm:"foreignKey:InterestID"`
	Sender   User     `gorm:"foreignKey:SenderID"`
	Receiver User     `gorm:"foreignKey:ReceiverID"`
}

// DB to Domain
func CommentDBToDomain(db Comment) comment.Comment {
	return comment.Comment{
		ID:         db.ID,
		InterestID: db.InterestID,
		SenderID:   db.SenderID,
		ReceiverID: db.ReceiverID,
		Content:    db.Content,
		IsRead:     db.IsRead,
		CreatedAt:  db.CreatedAt,
	}
}
