package usergooddeed

import "context"

type Repository interface {
	CreateGoodDeed(ctx context.Context, goodDeed *UserGoodDeed) error
	DeleteGoodDeed(ctx context.Context, transactionID uint, userID uint) error
}
