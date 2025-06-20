package transaction

import (
	"final_project/internal/pkg/enums"
	"time"
)

type Transaction struct {
	ID         uint
	PostID     uint
	InterestID uint
	SenderID   uint
	ReceiverID uint
	Status     int
	Method     string
	Items      []TransactionItem
}

type TransactionItem struct {
	TransactionID uint
	PostItemID    uint
	Quantity      int
}

type FilterTransaction struct {
	Page        int
	Limit       int
	Sort        string
	Order       string
	PostID      uint
	Status      enums.TransactionStatus
	SearchBy    string
	SearchValue string
}

type DetailTransaction struct {
	ID           uint
	InterestID   uint
	SenderID     uint
	ReceiverID   uint
	SenderName   string
	ReceiverName string
	Status       int
	Method       string
	Items        []DetailTransactionItem
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type DetailTransactionItem struct {
	ItemID     uint
	ItemName   string
	ItemImage  string
	PostItemID uint
	Quantity   int
}
