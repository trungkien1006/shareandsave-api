package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/transaction"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"final_project/internal/pkg/enums"
	"strconv"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepoDB struct {
	db *gorm.DB
}

func NewTransactionRepoDB(db *gorm.DB) *TransactionRepoDB {
	return &TransactionRepoDB{db: db}
}

func (r *TransactionRepoDB) GetAll(ctx context.Context, transactions *[]transaction.DetailTransaction, req transaction.FilterTransaction) (int, error) {
	var (
		query         *gorm.DB
		totalRecords  int64
		dbTransaction []dbmodel.Transaction
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Transaction{}).
		Table("transaction").
		Preload("Sender").
		Preload("Receiver").
		Preload("TransactionItems").
		Preload("TransactionItems.PostItem").
		Preload("TransactionItems.PostItem.Item").
		Joins("JOIN user as sender ON sender.id = transaction.sender_id").
		Joins("JOIN user as receiver ON receiver.id = transaction.receiver_id")

	if req.SearchBy != "" && req.SearchValue != "" {
		column := strcase.ToSnake(req.SearchBy) // "fullName" -> "full_name"

		if column == "sender_name" {
			column = "sender.full_name"
		} else if column == "receiver_name" {
			column = "receiver.full_name"
		} else {
			column = "transaction." + column
		}

		query.Where(column+" LIKE ? ", "%"+req.SearchValue+"%")
	}

	if req.Status != 0 {
		query.Where("transaction.status = ? ", req.Status)
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, err
	}

	if req.Sort != "" && req.Order != "" {
		query = query.Order("transaction." + strcase.ToSnake(req.Sort) + " " + req.Order)
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&dbTransaction).Error; err != nil {
		return 0, err
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	for _, value := range dbTransaction {
		*transactions = append(*transactions, dbmodel.TransactionDBToDetailDomain(value))
	}

	return totalPages, nil
}

func (r *TransactionRepoDB) GetByID(ctx context.Context, transactionID uint, transaction *transaction.Transaction) error {
	var dbTransaction dbmodel.Transaction

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Transaction{}).Where("id = ?", transactionID).First(&dbTransaction).Error; err != nil {
		return errors.New("Có lỗi khi truy vấn giao dịch theo id: " + err.Error())
	}

	*transaction = dbmodel.TransactionDBToDomain(dbTransaction)

	return nil
}

func (r *TransactionRepoDB) Create(ctx context.Context, transaction *transaction.Transaction) error {
	var (
		dbTransaction dbmodel.Transaction
		postID        uint
		senderID      uint
	)

	dbTransaction = dbmodel.TransactionDomainToDB(*transaction)

	tx := r.db.Debug().Begin()

	// Kiểm tra món đồ có tồn tại hay không
	for _, value := range transaction.Items {
		var postItem dbmodel.PostItem

		if err := tx.WithContext(ctx).
			Model(&dbmodel.PostItem{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", value.PostItemID).
			First(&postItem).Error; err != nil {
			tx.Rollback()
			return errors.New("Có lỗi khi kiểm tra đồ trong bài viết tồn tại: " + err.Error())
		}

		if value.Quantity > postItem.CurrentQuantity {
			tx.Rollback()
			return errors.New("Món đồ giao dịch không được có số lượng lớn hơn cho phép: id món đồ " + strconv.Itoa(int(postItem.ItemID)))
		}
	}

	// Lấy ra id của post theo interest id
	if err := tx.WithContext(ctx).Model(&dbmodel.Interest{}).
		Select("post_id").Where("id = ?", transaction.InterestID).
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

	// // Cập nhật lại số lượng đồ đạc ở post_item
	// for _, value := range dbTransaction.TransactionItems {
	// 	if err := tx.WithContext(ctx).Model(&dbmodel.PostItem{}).
	// 		Where("id = ?", value.PostItemID).
	// 		Update("current_quantity", gorm.Expr("current_quantity - ?", value.Quantity)).Error; err != nil {
	// 		tx.Rollback()
	// 		return errors.New("Có lỗi khi cập nhật lại số lượng đồ ở bài viết: " + err.Error())
	// 	}
	// }

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi commit transaction: " + err.Error())
	}

	*transaction = dbmodel.TransactionDBToDomain(dbTransaction)

	return nil
}

func (r *TransactionRepoDB) Update(ctx context.Context, transaction *transaction.Transaction) error {
	var (
		dbTransaction dbmodel.Transaction
	)

	dbTransaction = dbmodel.TransactionDomainToDB(*transaction)

	tx := r.db.Debug().Begin()

	// Kiểm tra món đồ có tồn tại hay không và số lượng so với cho phép trong bài viết
	for _, value := range transaction.Items {
		var postItem dbmodel.PostItem

		if err := tx.WithContext(ctx).
			Model(&dbmodel.PostItem{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", value.PostItemID).
			First(&postItem).Error; err != nil {
			tx.Rollback()
			return errors.New("Có lỗi khi kiểm tra số lượng đồ trong giao dịch: " + err.Error())
		}

		if value.Quantity > postItem.CurrentQuantity {
			tx.Rollback()
			return errors.New("Món đồ giao dịch không được có số lượng lớn hơn cho phép: id món đồ " + strconv.Itoa(int(postItem.ItemID)))
		}
	}

	// Cập nhật giao dịch
	if err := tx.WithContext(ctx).Model(&dbmodel.Transaction{}).
		Where("id = ?", dbTransaction.ID).
		Updates(&dbTransaction).Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi cập nhật giao dịch: " + err.Error())
	}

	if dbTransaction.Status == int(enums.TransactionStatusSuccess) {
		// Cập nhật lại số lượng đồ đạc ở post_item
		for _, value := range dbTransaction.TransactionItems {
			if err := tx.WithContext(ctx).Model(&dbmodel.PostItem{}).
				Where("id = ?", value.PostItemID).
				Update("current_quantity", gorm.Expr("current_quantity - ?", value.Quantity)).Error; err != nil {
				tx.Rollback()
				return errors.New("Có lỗi khi cập nhật lại số lượng đồ ở bài viết: " + err.Error())
			}
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

func (r *TransactionRepoDB) IsExist(ctx context.Context, transactionID uint) (bool, error) {
	var count int64

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Transaction{}).
		Where("id = ?", transactionID).
		Count(&count).Error; err != nil {
		return false, errors.New("Có lỗi khi kiểm tra giao dịch tồn tại: " + err.Error())
	}

	return count > 0, nil
}
