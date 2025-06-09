package transactiondto

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
