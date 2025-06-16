package comment

import "time"

type GetComment struct {
	InterestID int
	Page       int
	Limit      int
	Search     string
}

type Comment struct {
	ID         uint
	InterestID uint
	SenderID   uint
	ReceiverID uint
	Content    string
	IsRead     uint
	CreatedAt  time.Time
}
