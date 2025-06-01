package postdto

import (
	"time"
)

type AdminPostDTO struct {
	ID         uint      `json:"id"`
	AuthorName string    `json:"authorName"`
	Type       int       `json:"type"`
	Title      string    `json:"title"`
	Status     int8      `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}

type PostDTO struct {
	ID         uint     `json:"id"`
	AuthorName string   `json:"authorName"`
	Content    string   `json:"content"`
	Slug       string   `json:"slug"`
	Type       int      `json:"type"`
	Title      string   `json:"title"`
	Status     int8     `json:"status"`
	Images     []string `json:"images"`
}

type DetailPostItemDTO struct {
	ItemID     uint   `json:"itemID"`
	CategoryID uint   `json:"categoryID"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
}

type InterestDTO struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"userID"`
	UserName   string `json:"userName"`
	UserAvatar string `json:"userAvatar"`
	PostID     uint   `json:"postID"`
	Status     int    `json:"status"`
}

type DetailPostDTO struct {
	ID         uint                `json:"id"`
	AuthorID   uint                `json:"authorID"`
	AuthorName string              `json:"authorName"`
	Type       int                 `json:"type"`
	Slug       string              `json:"slug"`
	Title      string              `json:"title"`
	Content    string              `json:"content"`
	Info       string              `json:"info"`
	Status     int8                `json:"status"`
	Images     []string            `json:"images"`
	CreatedAt  time.Time           `json:"createdAt"`
	Tag        []string            `json:"tags"`
	Interest   []InterestDTO       `json:"interests"`
	Items      []DetailPostItemDTO `json:"items"`
}
