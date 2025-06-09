package transactiondto

import "final_project/internal/pkg/enums"

type CreateTransactionRequest struct {
	// PostID     uint                    `json:"postID" binding:"required"`
	InterestID uint `json:"interestID" binding:"required"`
	// SenderID   uint                    `json:"senderID" binding:"required"`
	Items []CreateTransactionItem `json:"items" binding:"required"`
}

type CreateTransactionItem struct {
	PostItemID uint `json:"postItemID" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required"`
}

type UpdateTransactionRequest struct {
	Status enums.TransactionStatus `json:"status"`
	Items  []UpdateTransactionItem `json:"items"`
}

type UpdateTransactionItem struct {
	TransactionID uint `json:"transactionID" binding:"required"`
	PostItemID    uint `json:"postItemID" binding:"required"`
	Quantity      int  `json:"quantity" binding:"required"`
}
