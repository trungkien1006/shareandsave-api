package usergooddeed

import (
	"time"
)

type UserGoodDeed struct {
	ID            uint
	UserID        uint
	UserName      string
	GoodDeedType  int
	GoodPoint     int
	TransactionID uint
	CreatedAt     time.Time
}
