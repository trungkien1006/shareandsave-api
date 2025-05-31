package post

import "time"

type Post struct {
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
}

type AdminPost struct {
	ID         uint
	AuthorName string
	Type       int
	Title      string
	Status     int8
	CreateAt   time.Time
}
