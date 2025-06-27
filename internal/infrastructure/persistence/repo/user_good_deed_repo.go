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
