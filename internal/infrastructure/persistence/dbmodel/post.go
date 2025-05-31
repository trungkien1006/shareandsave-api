package dbmodel

import (
	"final_project/internal/domain/post"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	AuthorID  uint `gorm:"index"`
	Type      int
	Slug      string `gorm:"unique;size:255"`
	Title     string `gorm:"size:255"`
	Content   string `gorm:"type:JSON"`
	Info      string `gorm:"type:JSON"`
	Status    int8   `gorm:"type:TINYINT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Author User `gorm:"foreignKey:AuthorID"`

	// 1-n: Một post có nhiều interest, post_item_warehouse
	Interests          []Interest          `gorm:"foreignKey:PostID"`
	PostItemWarehouses []PostItemWarehouse `gorm:"foreignKey:PostID"`
}

// DB → Domain
func PostDBToAdminPostDomain(dbPost Post) post.AdminPost {
	return post.AdminPost{
		ID:         dbPost.ID,
		AuthorName: dbPost.Author.FullName,
		Type:       dbPost.Type,
		Title:      dbPost.Title,
		Status:     dbPost.Status,
		CreateAt:   dbPost.CreatedAt,
	}
}

// Domain → DB
func PostDomainToDB(domainPost post.Post) Post {
	return Post{
		ID:       domainPost.ID,
		AuthorID: domainPost.AuthorID,
		Type:     domainPost.Type,
		Slug:     domainPost.Slug,
		Title:    domainPost.Title,
		Content:  domainPost.Content,
		Info:     domainPost.Info,
		Status:   domainPost.Status,
	}
}

// Db -> Domain
func PostDBToDomain(db Post) post.Post {
	return post.Post{
		ID:         db.ID,
		AuthorID:   db.AuthorID,
		AuthorName: db.Author.FullName,
		Type:       db.Type,
		Slug:       db.Slug,
		Title:      db.Title,
		Content:    db.Content,
		Info:       db.Info,
		Status:     db.Status,
	}
}
