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
	IsInterest bool      `json:"isInterest"`
}

type PostDTO struct {
	ID           uint     `json:"id"`
	AuthorName   string   `json:"authorName"`
	AuthorAvatar string   `json:"authorAvatar"`
	Content      string   `json:"content"`
	Slug         string   `json:"slug"`
	Type         int      `json:"type"`
	Title        string   `json:"title"`
	Status       int8     `json:"status"`
	Images       []string `json:"images"`
}

type DetailPostItemDTO struct {
	ID              uint   `json:"id"`
	ItemID          uint   `json:"itemID"`
	CategoryID      uint   `json:"categoryID"`
	CategoryName    string `json:"categoryName"`
	Image           string `json:"image"`
	Name            string `json:"name"`
	Quantity        int    `json:"quantity"`
	CurrentQuantity int    `json:"currentQuantity"`
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
	ID           uint                `json:"id"`
	AuthorID     uint                `json:"authorID"`
	AuthorName   string              `json:"authorName"`
	AuthorAvatar string              `json:"authorAvatar"`
	Type         int                 `json:"type"`
	Slug         string              `json:"slug"`
	Title        string              `json:"title"`
	Description  string              `json:"description"`
	Content      string              `json:"content"`
	Info         string              `json:"info"`
	Status       int8                `json:"status"`
	Images       []string            `json:"images"`
	CreatedAt    time.Time           `json:"createdAt"`
	Tag          []string            `json:"tags"`
	Interest     []InterestDTO       `json:"interests"`
	Items        []DetailPostItemDTO `json:"items"`
}

type PostWithCountDTO struct {
	ID               uint      `json:"id"`
	AuthorID         uint      `json:"authorID"`
	AuthorName       string    `json:"authorName"`
	AuthorAvatar     string    `json:"authorAvatar"`
	Type             int       `json:"type"`
	Slug             string    `json:"slug"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Content          string    `json:"content"`
	Info             string    `json:"info"`
	Status           int8      `json:"status"`
	Images           []string  `json:"images"`
	CreatedAt        time.Time `json:"createdAt"`
	Tag              []string  `json:"tags"`
	InterestCount    uint      `json:"interestCount"`
	ItemCount        uint      `json:"itemCount"`
	IsInterest       int       `json:"isInterest"`
	CurrentItemCount uint      `json:"currentItemCount"`
}
