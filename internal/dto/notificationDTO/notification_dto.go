package notificationdto

import "time"

type NotificationDTO struct {
	ID         uint      `json:"id"`
	SenderID   *uint     `json:"senderID"`
	ReceiverID uint      `json:"receiverID"`
	Type       string    `json:"type"`
	TargetType string    `json:"targetType"`
	TargetID   uint      `json:"targetID"`
	Content    string    `json:"content"`
	IsRead     bool      `json:"isRead"`
	CreatedAt  time.Time `json:"createdAt"`
}
