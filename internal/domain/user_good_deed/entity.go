package usergooddeed

import (
	"final_project/internal/domain/transaction"
	"time"
)

type UserGoodDeed struct {
	ID            uint
	UserID        uint
	UserName      string
	GoodDeedType  int
	GoodPoint     int
	TransactionID *uint
	CreatedAt     time.Time
}

type UserGoodDeedDetail struct {
	ID            uint
	UserID        uint
	UserName      string
	GoodDeedType  int
	GoodPoint     int
	TransactionID *uint
	CreatedAt     time.Time
	Items         []transaction.DetailTransactionItem
}
