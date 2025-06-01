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
