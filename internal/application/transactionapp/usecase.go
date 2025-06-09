package transactionapp

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/domain/item"
	"final_project/internal/domain/post"
	"final_project/internal/domain/transaction"
	"final_project/internal/domain/user"
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
	// Kiểm tra người sẽ tặng đồ có tồn tại hay không
	senderExisted, err := uc.userRepo.IsExist(ctx, transaction.SenderID)
	if err != nil {
		return err
	}

	if !senderExisted {
		return errors.New("Người chủ bài viết bạn muốn tạo giao dịch, không tồn tại")
	}

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
		itemExisted, err := uc.itemRepo.IsExist(ctx, value.ItemID)
		if err != nil {
			return err
		}

		if !itemExisted {
			return errors.New("Món đồ không tồn tại")
		}

		err = uc.postRepo.CheckPostItemQuantityOver(ctx, transaction.PostID, value.ItemID, value.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}
