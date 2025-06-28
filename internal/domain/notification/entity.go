package notification

import "time"

type Notification struct {
	ID         uint
	SenderID   uint
	ReceiverID uint
	Type       string
	TargetType string
	TargetID   uint
	Content    string
	IsRead     bool
	CreatedAt  time.Time
}
