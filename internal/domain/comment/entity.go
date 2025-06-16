package comment

import "time"

type GetComment struct {
	Page  uint
	Limit uint
}

type Comment struct {
	ID           uint
	InterestID   uint
	SenderID     uint
	SenderName   string
	ReceiverID   uint
	ReceiverName uint
	Content      string
	IsRead       uint
	CreatedAt    time.Time
}
