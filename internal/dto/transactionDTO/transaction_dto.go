package transactiondto

type TransactionDTO struct {
	ID         uint                 `json:"id"`
	PostID     uint                 `json:"postID"`
	InterestID uint                 `json:"interestID"`
	SenderID   uint                 `json:"senderID"`
	ReceiverID uint                 `json:"receiverID"`
	Items      []TransactionItemDTO `json:"items"`
}

type TransactionItemDTO struct {
	ItemID   uint `json:"itemID"`
	Quantity int  `json:"quantity"`
}
