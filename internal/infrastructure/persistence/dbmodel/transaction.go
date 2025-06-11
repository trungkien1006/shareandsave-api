package dbmodel

import (
	"final_project/internal/domain/transaction"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	InterestID uint `gorm:"index"`
	SenderID   uint `gorm:"index"`
	ReceiverID uint `gorm:"index"`
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	Interest Interest `gorm:"foreignKey:InterestID"`
	Sender   User     `gorm:"foreignKey:SenderID"`
	Receiver User     `gorm:"foreignKey:ReceiverID"`

	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionID"`
}

// Domain to DB
func TransactionDomainToDB(domain transaction.Transaction) Transaction {
	var dbItem []TransactionItem

	for _, value := range domain.Items {
		dbItem = append(dbItem, TransactionItem{
			PostItemID: value.PostItemID,
			Quantity:   value.Quantity,
		})
	}

	return Transaction{
		ID:               domain.ID,
		InterestID:       domain.InterestID,
		SenderID:         domain.SenderID,
		ReceiverID:       domain.ReceiverID,
		TransactionItems: dbItem,
		Status:           domain.Status,
	}
}

// DB to Domain
func TransactionDBToDomain(db Transaction) transaction.Transaction {
	var dbItem []transaction.TransactionItem

	for _, value := range db.TransactionItems {
		dbItem = append(dbItem, transaction.TransactionItem{
			TransactionID: value.TransactionID,
			PostItemID:    value.PostItemID,
			Quantity:      value.Quantity,
		})
	}

	return transaction.Transaction{
		ID:         db.ID,
		InterestID: db.InterestID,
		SenderID:   db.SenderID,
		ReceiverID: db.ReceiverID,
		Items:      dbItem,
		Status:     db.Status,
	}
}

// DB to Domain
func TransactionDBToDetailDomain(db Transaction) transaction.DetailTransaction {
	var domainItems []transaction.DetailTransactionItem

	for _, value := range db.TransactionItems {
		domainItems = append(domainItems, transaction.DetailTransactionItem{
			ItemID:     value.PostItem.ItemID,
			ItemName:   value.PostItem.Item.Name,
			ItemImage:  value.PostItem.Image,
			PostItemID: value.PostItemID,
			Quantity:   value.Quantity,
		})
	}

	return transaction.DetailTransaction{
		ID:           db.ID,
		InterestID:   db.InterestID,
		SenderID:     db.SenderID,
		ReceiverID:   db.ReceiverID,
		SenderName:   db.Sender.FullName,
		ReceiverName: db.Receiver.FullName,
		Items:        domainItems,
		Status:       db.Status,
		CreatedAt:    db.CreatedAt,
		UpdatedAt:    db.UpdatedAt,
	}
}
