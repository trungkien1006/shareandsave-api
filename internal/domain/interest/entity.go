package interest

import "time"

type Interest struct {
	ID         uint
	UserID     uint
	UserName   string
	UserAvatar string
	PostID     uint
	Status     int
	CreatedAt  time.Time
}

type PostInterestItem struct {
	ID           uint
	Name         string
	CategoryName string
	Quantity     int
	Image        string
}

type PostInterest struct {
	ID          uint
	AuthorID    uint
	AuthorName  string
	Title       string
	Description string
	Slug        string
	Type        int
	Items       []PostInterestItem
	Interests   []Interest
	UpdatedAt   time.Time
}

type GetInterest struct {
	Page   int
	Limit  int
	Type   int
	Sort   string
	Order  string
	Search string
}

type CreateInterest struct {
	UserID uint
	PostID uint
}
