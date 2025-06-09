package transaction

type Transaction struct {
	ID         uint
	PostID     uint
	InterestID uint
	SenderID   uint
	ReceiverID uint
	Status     int
	Items      []TransactionItem
}

type TransactionItem struct {
	TransactionID uint
	ItemID        uint
	Quantity      int
}
