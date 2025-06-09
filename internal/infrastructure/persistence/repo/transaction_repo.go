package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/transaction"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type TransactionRepoDB struct {
	db *gorm.DB
}

func NewTransactionRepoDB(db *gorm.DB) *TransactionRepoDB {
	return &TransactionRepoDB{db: db}
}

func (r *TransactionRepoDB) Create(ctx context.Context, transaction *transaction.Transaction) error {
	var (
		dbTransaction dbmodel.Transaction
		postID        uint
		senderID      uint
	)

	dbTransaction = dbmodel.TransactionDomainToDB(*transaction)

	tx := r.db.Debug().Begin()

	// Lấy ra id của post theo interest id
	if err := tx.WithContext(ctx).Model(&dbmodel.Interest{}).
		Select("post_id").Where("interest_id = ?", transaction.InterestID).
		Scan(&postID).Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi lấy post id từ interest id: " + err.Error())
	}

	// Lấy ra id của author theo post id
	if err := tx.WithContext(ctx).Model(&dbmodel.Post{}).
		Select("author_id").Where("id = ?", postID).
		Scan(&senderID).Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi lấy id của author theo post id: " + err.Error())
	}

	dbTransaction.SenderID = senderID

	// Tạo giao dịch
	if err := tx.WithContext(ctx).Model(&dbmodel.Transaction{}).
		Create(&dbTransaction).Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi tạo giao dịch mới: " + err.Error())
	}

	// Cập nhật lại số lượng đồ đạc ở post_item
	for _, value := range dbTransaction.TransactionItems {
		if err := tx.WithContext(ctx).Model(&dbmodel.PostItem{}).
			Where("id = ?", value.PostItemID).
			Update("quantity", gorm.Expr("quantity - ?", value.Quantity)).Error; err != nil {
			tx.Rollback()
			return errors.New("Có lỗi khi cập nhật lại số lượng đồ ở bài viết: " + err.Error())
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi commit transaction: " + err.Error())
	}

	*transaction = dbmodel.TransactionDBToDomain(dbTransaction)

	return nil
}
