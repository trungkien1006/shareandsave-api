package post

import (
	"final_project/internal/domain/interest"
	"final_project/internal/domain/item"
	"time"
)

type OldItemsInPost struct {
	ItemID   uint
	Quantity int
}

type NewItemsInPost struct {
	ItemID     uint
	CategoryID uint
	Image      string
	Name       string
	Quantity   int
}

type CreatePost struct {
	ID          uint
	AuthorID    uint
	AuthorName  string
	Type        int
	Slug        string
	Title       string
	Content     string
	Info        string
	Description string
	Status      int8
	Images      []string
	Tag         []string
	Items       []item.Item
	OldItems    []OldItemsInPost
	NewItems    []NewItemsInPost
}

type Post struct {
	ID         uint
	AuthorID   uint
	AuthorName string
	Type       int
	Slug       string
	Title      string
	Content    string
	Info       string
	Status     int8
	Images     []string
	CreatedAt  time.Time
}

type DetailPostItem struct {
	ItemID     uint
	CategoryID uint
	Image      string
	Name       string
	Quantity   int
}

type DetailPost struct {
	ID          uint
	AuthorID    uint
	AuthorName  string
	Type        int
	Slug        string
	Title       string
	Description string
	Content     string
	Info        string
	Status      int8
	Images      []string
	CreatedAt   time.Time
	Tag         []string
	Interest    []interest.Interest
	Items       []DetailPostItem
}

type PostFilterRequest struct {
	Page        int
	Limit       int
	Sort        string
	Order       string
	Status      int
	Type        int
	SearchBy    string
	SearchValue string
}
