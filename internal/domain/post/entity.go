package post

import (
	"final_project/internal/domain/item"
	"os/user"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	AuthorID    uint   `gorm:"index"`
	ItemID      uint   `gorm:"index"`
	Title       string `gorm:"size:255"`
	Description string `gorm:"type:TEXT"`
	Status      int8   `gorm:"type:TINYINT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Author user.User `gorm:"foreignKey:AuthorID"`
	Item   item.Item `gorm:"foreignKey:ItemID"`
}

func NewPost(authorID uint, itemID uint, title, description string, status int8) *Post {
	return &Post{
		AuthorID:    authorID,
		ItemID:      itemID,
		Title:       title,
		Description: description,
		Status:      status,
	}
}
