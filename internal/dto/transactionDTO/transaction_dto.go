package transactiondto

import "time"

type TransactionDTO struct {
	ID         uint                 `json:"id"`
	InterestID uint                 `json:"interestID"`
	SenderID   uint                 `json:"senderID"`
	ReceiverID uint                 `json:"receiverID"`
	Status     int                  `json:"status"`
	Items      []TransactionItemDTO `json:"items"`
}

type TransactionItemDTO struct {
	PostItemID uint `json:"postItemID"`
	Quantity   int  `json:"quantity"`
}

type DetailTransactionDTO struct {
	ID           uint                       `json:"id"`
	InterestID   uint                       `json:"interestID"`
	SenderID     uint                       `json:"senderID"`
	ReceiverID   uint                       `json:"receiverID"`
	SenderName   string                     `json:"senderName"`
	ReceiverName string                     `json:"receiverName"`
	Status       int                        `json:"status"`
	Items        []DetailTransactionItemDTO `json:"items"`
	CreatedAt    time.Time                  `json:"createdAt"`
	UpdatedAt    time.Time                  `json:"updatedAt"`
}

type DetailTransactionItemDTO struct {
	ItemID     uint   `json:"itemID"`
	ItemName   string `json:"itemName"`
	ItemImage  string `json:"itemImage"`
	PostItemID uint   `json:"postItemID"`
	Quantity   int    `json:"quantity"`
}
