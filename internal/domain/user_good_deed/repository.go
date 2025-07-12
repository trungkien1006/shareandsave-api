package usergooddeed

import (
	"context"
	"final_project/internal/domain/user"
)

type Repository interface {
	GetUserReport(ctx context.Context, userReports *[]user.UserReport) error
	GetUserGoodDeed(ctx context.Context, userGoodDeeds *[]UserGoodDeedDetail, userID int) error
	GetByID(ctx context.Context, goodDeed *UserGoodDeed, goodDeedID uint) error
	CreateGoodDeed(ctx context.Context, goodDeed *UserGoodDeed) error
	DeleteGoodDeed(ctx context.Context, transactionID uint, userID uint) error
	DeleteGoodDeedByID(ctx context.Context, goodDeedID uint) error
}
