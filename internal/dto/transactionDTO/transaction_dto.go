package transactiondto

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
