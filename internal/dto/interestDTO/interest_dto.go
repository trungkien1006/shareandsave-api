package interestdto

import "final_project/internal/pkg/enums"

type Interest struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"userID"`
	UserName   string `json:"userName"`
	UserAvatar string `json:"userAvatar"`
	PostID     uint   `json:"postID"`
	Status     int    `json:"status"`
}

type PostInterestItem struct {
	ID           uint   `json:"id`
	Name         string `json:"name"`
	CategoryName string `json:"categoryName"`
	Quantity     int    `json:"quantity"`
	Image        string `json:"image"`
}

type PostInterest struct {
	ID        uint               `json:"id"`
	Title     string             `json:"title"`
	Type      enums.InterestType `json:"type"`
	Items     []PostInterestItem `json:"items"`
	Interests []Interest         `json:"interests"`
}
