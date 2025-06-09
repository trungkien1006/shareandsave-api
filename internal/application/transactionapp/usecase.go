package transactionapp

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/domain/item"
	"final_project/internal/domain/post"
	"final_project/internal/domain/transaction"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
)

type UseCase struct {
	repo         transaction.Repository
	userRepo     user.Repository
	interestRepo interest.Repository
	itemRepo     item.Repository
	postRepo     post.Repository
}

func NewUseCase(r transaction.Repository, userRepo user.Repository, interestRepo interest.Repository, itemRepo item.Repository, postRepo post.Repository) *UseCase {
	return &UseCase{
		repo:         r,
		userRepo:     userRepo,
		interestRepo: interestRepo,
		itemRepo:     itemRepo,
		postRepo:     postRepo,
	}
}

func (uc *UseCase) CreateTransaction(ctx context.Context, transaction *transaction.Transaction) error {
	// Kiểm tra phiếu quan tâm có tồn tại hay không
	interestExisted, err := uc.interestRepo.IsExistByID(ctx, transaction.InterestID)
	if err != nil {
		return err
	}

	if !interestExisted {
		return errors.New("Quan tâm không tồn tại hoặc đã bị xóa")
	}

	// Kiểm tra món đồ có tồn tại hay không và số lượng so với cho phép trong bài viết
	for _, value := range transaction.Items {
		err = uc.postRepo.CheckPostItemQuantityOver(ctx, value.PostItemID, value.Quantity)
		if err != nil {
			return err
		}
	}

	transaction.Status = int(enums.TransactionStatusPending)

	if err := uc.repo.Create(ctx, transaction); err != nil {
		return err
	}

	return nil
}
