package post

import "time"

type CreatePost struct {
	ID          uint
	AuthorID    uint
	AuthorName  string
	FullName    string
	Email       string
	PhoneNumber string
	Type        int
	Slug        string
	Title       string
	Content     string
	Info        string
	Status      int8
	Images      []string
	Tag         []string
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
