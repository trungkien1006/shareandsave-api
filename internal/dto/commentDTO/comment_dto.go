package commentdto

import "time"

type Comment struct {
	ID         uint      `json:"id"`
	InterestID uint      `json:"interestID"`
	SenderID   uint      `json:"senderID"`
	ReceiverID uint      `json:"receiverID"`
	Content    string    `json:"content"`
	IsRead     uint      `json:"isread"`
	CreatedAt  time.Time `json:"createdAt"`
}
