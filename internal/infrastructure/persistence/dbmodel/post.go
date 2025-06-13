package dbmodel

import (
	"final_project/internal/domain/interest"
	"final_project/internal/domain/post"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	AuthorID    uint `gorm:"index"`
	Type        int
	Slug        string      `gorm:"unique;size:255"`
	Title       string      `gorm:"size:255"`
	Description string      `gorm:"type:MEDIUMTEXT"`
	Content     string      `gorm:"type:JSON"`
	Info        string      `gorm:"type:JSON"`
	Status      int8        `gorm:"type:TINYINT"`
	Image       StringArray `gorm:"type:JSON"`
	Tag         StringArray `gorm:"type:JSON"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Author User `gorm:"foreignKey:AuthorID"`

	// 1-n: Một post có nhiều interest, post_item_warehouse
	Interests []Interest `gorm:"foreignKey:PostID"`
	PostItem  []PostItem `gorm:"foreignKey:PostID"`
}

type AdminPost struct {
	Post
	IsInterest   bool   `gorm:"column:is_interested"`
	AuthorAvatar string `gorm:"column:"author_avatar"`
	AuthorName   string `gorm:"column:author_name"`
}

type PostWithCounts struct {
	Post                    // Lấy toàn bộ trường từ bảng post
	AuthorAvatar     string `gorm:"column:"author_avatar"`
	AuthorName       string `gorm:"column:author_name"`
	InterestCount    int64  `gorm:"column:interest_count"`
	ItemCount        int64  `gorm:"column:item_count"`
	CurrentItemCount int    `gorm:"column:current_item_count"`
}

// DB → Domain
func AdminPostDBToDomain(dbPost AdminPost) post.Post {

	return post.Post{
		ID:           dbPost.ID,
		AuthorName:   dbPost.AuthorName,
		AuthorAvatar: dbPost.AuthorAvatar,
		Type:         dbPost.Type,
		Title:        dbPost.Title,
		Status:       dbPost.Status,
		CreatedAt:    dbPost.CreatedAt,
		Images:       dbPost.Image,
		IsInterested: dbPost.IsInterest,
	}
}

// DB → Domain
func PostDBToDomain(dbPost Post) post.Post {

	return post.Post{
		ID:           dbPost.ID,
		AuthorName:   dbPost.Author.FullName,
		AuthorAvatar: dbPost.Author.Avatar,
		Type:         dbPost.Type,
		Title:        dbPost.Title,
		Status:       dbPost.Status,
		CreatedAt:    dbPost.CreatedAt,
		Images:       dbPost.Image,
	}
}

// DB -> Domain
func PostWithCountDBToDomain(db PostWithCounts) post.PostWithCount {
	domainTag := make([]string, 0)
	domainImage := make([]string, 0)
	currentItemCount := 0
	itemCount := 0

	for _, value := range db.Tag {
		domainTag = append(domainTag, value)
	}

	for _, value := range db.Image {
		domainImage = append(domainImage, value)
	}

	for _, value := range db.PostItem {
		currentItemCount += value.CurrentQuantity
		itemCount += value.Quantity
	}

	return post.PostWithCount{
		ID:               db.ID,
		AuthorID:         db.AuthorID,
		AuthorName:       db.AuthorName,
		AuthorAvatar:     db.AuthorAvatar,
		Type:             db.Type,
		Slug:             db.Slug,
		Title:            db.Title,
		Description:      db.Description,
		Content:          db.Content,
		Info:             db.Info,
		Status:           db.Status,
		Images:           domainImage,
		CreatedAt:        db.CreatedAt,
		Tag:              domainTag,
		InterestCount:    uint(len(db.Interests)),
		ItemCount:        uint(itemCount),
		CurrentItemCount: uint(currentItemCount),
	}
}

// Domain → DB
func DetailPostDBToDomain(db Post) post.DetailPost {
	domainInterest := make([]interest.Interest, 0)
	domainPostItem := make([]post.DetailPostItem, 0)
	domainTag := make([]string, 0)
	domainImage := make([]string, 0)

	for _, value := range db.Interests {
		domainInterest = append(domainInterest, interest.Interest{
			ID:         value.ID,
			UserID:     value.UserID,
			UserName:   value.User.FullName,
			UserAvatar: value.User.Avatar,
			PostID:     value.PostID,
			Status:     value.Status,
		})
	}

	for _, value := range db.PostItem {
		domainPostItem = append(domainPostItem, post.DetailPostItem{
			ID:              value.ID,
			ItemID:          value.ItemID,
			CategoryID:      value.Item.CategoryID,
			CategoryName:    value.Item.Category.Name,
			Image:           value.Image,
			Name:            value.Item.Name,
			Quantity:        value.Quantity,
			CurrentQuantity: value.CurrentQuantity,
		})
	}

	for _, value := range db.Tag {
		domainTag = append(domainTag, value)
	}

	for _, value := range db.Image {
		domainImage = append(domainImage, value)
	}

	return post.DetailPost{
		ID:           db.ID,
		AuthorID:     db.AuthorID,
		AuthorName:   db.Author.FullName,
		AuthorAvatar: db.Author.Avatar,
		Type:         db.Type,
		Slug:         db.Slug,
		Title:        db.Title,
		Description:  db.Description,
		Content:      db.Content,
		Info:         db.Info,
		Status:       db.Status,
		Images:       domainImage,
		CreatedAt:    db.CreatedAt,
		Tag:          domainTag,
		Interest:     domainInterest,
		Items:        domainPostItem,
	}
}

// Domain → DB
func AdminPostDomainToDB(domainPost post.Post) Post {
	return Post{
		ID:          domainPost.ID,
		AuthorID:    domainPost.AuthorID,
		Type:        domainPost.Type,
		Slug:        domainPost.Slug,
		Title:       domainPost.Title,
		Content:     domainPost.Content,
		Description: domainPost.Description,
		Info:        domainPost.Info,
		Status:      domainPost.Status,
		Image:       domainPost.Images,
	}
}

// Domain → DB
func CreatePostDomainToDB(domainPost post.CreatePost) Post {
	var postItems []PostItem

	for _, value := range domainPost.OldItems {
		postItems = append(postItems, PostItem{
			ItemID:          value.ItemID,
			Image:           value.Image,
			Quantity:        value.Quantity,
			CurrentQuantity: value.Quantity,
		})
	}

	for _, value := range domainPost.NewItems {
		postItems = append(postItems, PostItem{
			ItemID:          value.ItemID,
			Image:           value.Image,
			Quantity:        value.Quantity,
			CurrentQuantity: value.Quantity,
		})
	}

	return Post{
		ID:          domainPost.ID,
		AuthorID:    domainPost.AuthorID,
		Type:        domainPost.Type,
		Slug:        domainPost.Slug,
		Title:       domainPost.Title,
		Description: domainPost.Description,
		Content:     domainPost.Content,
		Info:        domainPost.Info,
		Status:      domainPost.Status,
		Image:       domainPost.Images,
		Tag:         domainPost.Tag,
		PostItem:    postItems,
	}
}

// Db -> Domain
func PostDBToCreatePostDomain(db Post) post.CreatePost {
	return post.CreatePost{
		ID:         db.ID,
		AuthorID:   db.AuthorID,
		AuthorName: db.Author.FullName,
		Type:       db.Type,
		Slug:       db.Slug,
		Title:      db.Title,
		Content:    db.Content,
		Info:       db.Info,
		Status:     db.Status,
		Images:     db.Image,
		Tag:        db.Tag,
	}
}
