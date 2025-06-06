package interest

type Interest struct {
	ID         uint
	UserID     uint
	UserName   string
	UserAvatar string
	PostID     uint
	Status     int
}

type PostInterestItem struct {
	ID           uint
	Name         string
	CategoryName string
	Quantity     int
	Image        string
}

type PostInterest struct {
	ID        uint
	Title     string
	Type      int
	Items     []PostInterestItem
	Interests []Interest
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
