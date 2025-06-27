package dbmodel

import (
	"final_project/internal/domain/transaction"
	usergooddeed "final_project/internal/domain/user_good_deed"
	"time"

	"gorm.io/gorm"
)

type UserGoodDeed struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	UserID        uint `gorm:"index"` // Có thể null
	GoodDeedType  int  `gorm:"type:int"`
	GoodPoint     int  `gorm:"type:int"` // Điểm tốt, có thể null
	TransactionID uint `gorm:"index"`    // Có thể null
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	// Quan hệ: n-1 với User và Transaction
	User        User        `gorm:"foreignKey:UserID"`
	Transaction Transaction `gorm:"foreignKey:TransactionID"`
}

// Domain to DB
func GoodDeedDomainToDB(domain usergooddeed.UserGoodDeed) UserGoodDeed {
	return UserGoodDeed{
		ID:            domain.ID,
		UserID:        domain.UserID,
		GoodDeedType:  domain.GoodDeedType,
		GoodPoint:     domain.GoodPoint,
		TransactionID: domain.TransactionID,
	}
}

// DB to Domain
func GoodDeedDBToDomain(db UserGoodDeed) usergooddeed.UserGoodDeedDetail {
	items := make([]transaction.DetailTransactionItem, 0)

	for _, item := range db.Transaction.TransactionItems {
		items = append(items, transaction.DetailTransactionItem{
			ItemID:     item.PostItem.Item.ID,
			ItemName:   item.PostItem.Item.Name,
			ItemImage:  item.PostItem.Item.Image,
			PostItemID: item.PostItem.ID,
			Quantity:   item.Quantity,
		})
	}

	return usergooddeed.UserGoodDeedDetail{
		ID:            db.ID,
		UserID:        db.UserID,
		UserName:      db.User.FullName,
		GoodDeedType:  db.GoodDeedType,
		GoodPoint:     db.GoodPoint,
		TransactionID: db.TransactionID,
		CreatedAt:     db.CreatedAt,
		Items:         items,
	}
}
