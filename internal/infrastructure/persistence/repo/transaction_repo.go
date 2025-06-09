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
	var dbTransaction dbmodel.Transaction

	dbTransaction = dbmodel.TransactionDomainToDB(*transaction)

	tx := r.db.Debug().Begin()

	// Tạo giao dịch
	if err := tx.WithContext(ctx).Model(&dbmodel.Transaction{}).
		Create(&dbTransaction).Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi tạo giao dịch mới: " + err.Error())
	}

	// Cập nhật lại số lượng đồ đạc ở post_item
	for _, value := range dbTransaction.TransactionItems {
		if err := tx.WithContext(ctx).Model(&dbmodel.PostItem{}).
			Where("post_id = ? AND item_id = ?", transaction.PostID, value.ItemID).
			Update("quantity", gorm.Expr("quantity - ?", value.Quantity)).Error; err != nil {
			tx.Rollback()
			return errors.New("Có lỗi khi cập nhật lại số lượng đồ ở bài viết: " + err.Error())
		}
	}

	*transaction = dbmodel.TransactionDBToDomain(dbTransaction)

	return nil
}
