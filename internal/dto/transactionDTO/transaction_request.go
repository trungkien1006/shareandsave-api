package transactiondto

type CreateTransactionRequest struct {
	ID         uint                    `json:"id" binding:"required"`
	InterestID uint                    `json:"interestID" binding:"required"`
	SenderID   uint                    `json:"senderID" binding:"required"`
	ReceiverID uint                    `json:"receiverID" binding:"required"`
	Items      []CreateTransactionItem `json:"items" binding:"required"`
}

type CreateTransactionItem struct {
	ItemID   uint `json:"itemID" binding:"required"`
	Quantity int  `json:"quantity" binding:"required"`
}
