package transactionapp

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/domain/transaction"
	"final_project/internal/domain/user"
)

type UseCase struct {
	repo         transaction.Repository
	userRepo     user.Repository
	interestRepo interest.Repository
}

func NewUseCase(r transaction.Repository, userRepo user.Repository, interestRepo interest.Repository) *UseCase {
	return &UseCase{
		repo:         r,
		userRepo:     userRepo,
		interestRepo: interestRepo,
	}
}

func (uc *UseCase) CreateTransaction(ctx context.Context, transaction *transaction.Transaction) error {
	// Kiểm tra người sẽ tặng đồ có tồn tại hay không
	senderExisted, err := uc.userRepo.IsExist(ctx, transaction.SenderID)
	if err != nil {
		return err
	}

	if !senderExisted {
		return errors.New("Người bên kia giao dịch bạn muốn tạo không tồn tại")
	}

	return nil
}
