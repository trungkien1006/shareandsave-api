package notification

import "time"

type Notification struct {
	ID           uint
	SenderID     *uint
	SenderName   string
	ReceiverID   *uint
	ReceiverName string
	Type         string
	TargetType   string
	TargetID     uint
	Content      string
	IsRead       bool
	CreatedAt    time.Time
}

type GetAllNotiRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
