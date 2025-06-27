package usergooddeed

import "context"

type Repository interface {
	GetUserGoodDeed(ctx context.Context, userGoodDeeds *[]UserGoodDeedDetail, userID int) error
	GetByID(ctx context.Context, goodDeed *UserGoodDeed, goodDeedID uint) error
	CreateGoodDeed(ctx context.Context, goodDeed *UserGoodDeed) error
	DeleteGoodDeed(ctx context.Context, transactionID uint, userID uint) error
	DeleteGoodDeedByID(ctx context.Context, goodDeedID uint) error
}
