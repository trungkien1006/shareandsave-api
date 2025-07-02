package notificationdto

import "time"

type NotificationDTO struct {
	ID           uint      `json:"id"`
	SenderID     *uint     `json:"senderID"`
	SenderName   string    `json:"senderName"`
	ReceiverID   *uint     `json:"receiverID"`
	ReceiverName string    `json:"receiverName"`
	Type         string    `json:"type"`
	TargetType   string    `json:"targetType"`
	TargetID     uint      `json:"targetID"`
	Content      string    `json:"content"`
	IsRead       bool      `json:"isRead"`
	CreatedAt    time.Time `json:"createdAt"`
}
