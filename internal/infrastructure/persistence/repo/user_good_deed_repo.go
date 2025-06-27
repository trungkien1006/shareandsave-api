package persistence

import (
	"context"
	"errors"
	usergooddeed "final_project/internal/domain/user_good_deed"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type UserGoodDeedRepoDB struct {
	db *gorm.DB
}

func NewUserGoodDeedRepoDB(db *gorm.DB) *UserGoodDeedRepoDB {
	return &UserGoodDeedRepoDB{db: db}
}

func (r *UserGoodDeedRepoDB) GetUserGoodDeed(ctx context.Context, userGoodDeeds *[]usergooddeed.UserGoodDeedDetail, userID int) error {
	var dbGoodDeeds []dbmodel.UserGoodDeed

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.UserGoodDeed{}).
		Where("user_id = ?", userID).
		Preload("User").
		Preload("Transaction").
		Preload("Transaction.TransactionItems").
		Preload("Transaction.TransactionItems.PostItem").
		Preload("Transaction.TransactionItems.PostItem.Item").
		Order("created_at DESC").
		Find(&dbGoodDeeds).Error; err != nil {
		return errors.New("Có lỗi khi lấy việc tốt của người dùng: " + err.Error())
	}

	for _, dbGoodDeed := range dbGoodDeeds {
		*userGoodDeeds = append(*userGoodDeeds, dbmodel.GoodDeedDBToDomain(dbGoodDeed))
	}

	return nil
}

func (r *UserGoodDeedRepoDB) CreateGoodDeed(ctx context.Context, goodDeed *usergooddeed.UserGoodDeed) error {
	var dbGoodDeed dbmodel.UserGoodDeed

	dbGoodDeed = dbmodel.GoodDeedDomainToDB(*goodDeed)

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.UserGoodDeed{}).Create(&dbGoodDeed).Error; err != nil {
		return errors.New("Có lỗi khi tạo việc tốt: " + err.Error())
	}

	return nil
}

func (r *UserGoodDeedRepoDB) DeleteGoodDeed(ctx context.Context, transactionID uint, userID uint) error {
	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.UserGoodDeed{}).Where("transaction_id = ? AND user_id = ?", transactionID, userID).Delete(&dbmodel.UserGoodDeed{}).Error; err != nil {
		return errors.New("Có lỗi khi xóa việc tốt: " + err.Error())
	}

	return nil
}
