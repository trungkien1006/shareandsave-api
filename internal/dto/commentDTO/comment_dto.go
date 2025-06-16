package commentdto

import "time"

type CommentDTO struct {
	ID         uint      `json:"id"`
	InterestID uint      `json:"interestID"`
	SenderID   uint      `json:"senderID"`
	ReceiverID uint      `json:"receiverID"`
	Content    string    `json:"message"`
	IsRead     uint      `json:"isRead"`
	CreatedAt  time.Time `json:"createdAt"`
}
