package interestdto

import (
	"final_project/internal/pkg/enums"
	"time"
)

type Interest struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"userID"`
	UserName   string    `json:"userName"`
	UserAvatar string    `json:"userAvatar"`
	PostID     uint      `json:"postID"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}

type PostInterestItem struct {
	ID              uint   `json:"id"`
	ItemID          uint   `json:"itemID"`
	Name            string `json:"name"`
	CategoryName    string `json:"categoryName"`
	Quantity        int    `json:"quantity"`
	CurrentQuantity int    `json:"currentQuantity"`
	Image           string `json:"image"`
}

type PostInterest struct {
	ID          uint               `json:"id"`
	AuthorID    uint               `json:"authorID"`
	AuthorName  string             `json:"authorName"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Slug        string             `json:"slug"`
	Type        enums.PostType     `json:"type"`
	Items       []PostInterestItem `json:"items"`
	Interests   []Interest         `json:"interests"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}
