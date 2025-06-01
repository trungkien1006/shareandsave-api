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
	ItemID     uint
	CategoryID uint
	Image      string
	Name       string
	Quantity   int
}

type InterestDTO struct {
	ID         uint
	UserID     uint
	UserName   string
	UserAvatar string
	PostID     uint
	Status     int
}

type DetailPostDTO struct {
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
	Tag        []string
	Interest   []InterestDTO
	Items      []DetailPostItemDTO
}
