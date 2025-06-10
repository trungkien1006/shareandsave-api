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
	"strconv"
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

func (uc *UseCase) GetAllTransaction(ctx context.Context, transactions *[]transaction.DetailTransaction, filter transaction.FilterTransaction) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, transactions, filter)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
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

	transaction.Status = int(enums.TransactionStatusPending)

	if err := uc.repo.Create(ctx, transaction); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdateTransaction(ctx context.Context, domainTransaction *transaction.Transaction) error {
	var updateTransaction transaction.Transaction

	// Kiểm tra nếu không truyền lên gì cả
	if domainTransaction.Status == 0 && domainTransaction.Items == nil {
		return errors.New("Không có trường nào để cập nhật")
	}

	// Kiểm tra transaction có tồn tại
	transactionExisted, err := uc.repo.IsExist(ctx, domainTransaction.ID)
	if err != nil {
		return err
	}

	if !transactionExisted {
		return errors.New("Giao dịch không tồn tại: id giao dịch bạn gửi là " + strconv.Itoa(int(domainTransaction.ID)))
	}

	// Truy vấn transaction cần update
	if err := uc.repo.GetByID(ctx, domainTransaction.ID, &updateTransaction); err != nil {
		return err
	}

	if domainTransaction.Status != 0 {
		updateTransaction.Status = domainTransaction.Status
	}

	if domainTransaction.Items != nil {
		updateTransaction.Items = domainTransaction.Items
	}

	// Cập nhật
	if err := uc.repo.Update(ctx, &updateTransaction); err != nil {
		return err
	}

	return nil
}
