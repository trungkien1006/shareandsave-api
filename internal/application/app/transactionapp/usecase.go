package transactionapp

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/domain/item"
	"final_project/internal/domain/post"
	"final_project/internal/domain/transaction"
	"final_project/internal/domain/user"
	usergooddeed "final_project/internal/domain/user_good_deed"
	"final_project/internal/pkg/enums"
	"strconv"
	"time"
)

type UseCase struct {
	repo             transaction.Repository
	userRepo         user.Repository
	interestRepo     interest.Repository
	itemRepo         item.Repository
	postRepo         post.Repository
	userGoodDeedRepo usergooddeed.Repository
}

func NewUseCase(r transaction.Repository, userRepo user.Repository, interestRepo interest.Repository, itemRepo item.Repository, postRepo post.Repository, userGoodDeedRepo usergooddeed.Repository) *UseCase {
	return &UseCase{
		repo:             r,
		userRepo:         userRepo,
		interestRepo:     interestRepo,
		itemRepo:         itemRepo,
		postRepo:         postRepo,
		userGoodDeedRepo: userGoodDeedRepo,
	}
}

func (uc *UseCase) GetAllTransaction(ctx context.Context, transactions *[]transaction.DetailTransaction, filter transaction.FilterTransaction) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, transactions, filter)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetDetailTransactionByInterestID(ctx context.Context, transaction *transaction.DetailTransaction, interestID uint) error {
	if err := uc.repo.GetDetailTransactionByInterestID(ctx, transaction, interestID); err != nil {
		return err
	}

	return nil
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
	if domainTransaction.Status == 0 && domainTransaction.Items == nil && domainTransaction.Method == "" {
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

	tempStatus := updateTransaction.Status

	if domainTransaction.Status != 0 {
		if tempStatus == domainTransaction.Status {
			return errors.New("Trạng thái giao dịch không thay đổi")
		} else if tempStatus == int(enums.TransactionStatusPending) && domainTransaction.Status == int(enums.TransactionStatusRollBack) {
			return errors.New("Không thể thể rollback giao dịch đang chờ")
		} else if tempStatus == int(enums.TransactionStatusSuccess) && domainTransaction.Status == int(enums.TransactionStatusCancelled) {
			return errors.New("Không thể hủy giao dịch đã thành công")
		} else if tempStatus == int(enums.TransactionStatusCancelled) {
			return errors.New("Không thể cập nhật giao dịch đã bị hủy")
		} else if tempStatus == int(enums.TransactionStatusSuccess) && domainTransaction.Status != int(enums.TransactionStatusRollBack) {
			return errors.New("Không thể hủy hoặc quay về chờ giao dịch đã thành công")
		} else if tempStatus == int(enums.TransactionStatusRollBack) {
			return errors.New("Không thể cập nhật giao dịch đã rollback")
		}

		updateTransaction.Status = domainTransaction.Status
	}

	if domainTransaction.Items != nil {
		updateTransaction.Items = domainTransaction.Items
	}

	if domainTransaction.Method != "" {
		updateTransaction.Method = domainTransaction.Method
	}

	// Cập nhật
	if err := uc.repo.Update(ctx, &updateTransaction); err != nil {
		return err
	}

	if domainTransaction.Status == int(enums.TransactionStatusRollBack) && tempStatus == int(enums.TransactionStatusSuccess) {
		var (
			user     user.User
			postType int64
		)

		postType, err = uc.postRepo.GetPostType(ctx, updateTransaction.InterestID)
		if err != nil {
			return err
		}

		user.ID = domainTransaction.SenderID

		if postType == int64(enums.PostTypeGiveAwayOldItem) {
			user.GoodPoint = -enums.GoodPointGiveOldItem
		} else if postType == int64(enums.PostTypeFoundItem) {
			user.GoodPoint = -enums.GoodPointGiveLoseItem
		} else if postType == int64(enums.PostTypeCampaign) {
			user.GoodPoint = -enums.GoodPointJoinCampaign
		}

		if err := uc.userRepo.Update(ctx, &user); err != nil {
			return err
		}

		//Xóa việc tốt cho user
		if err := uc.userGoodDeedRepo.DeleteGoodDeed(ctx, domainTransaction.ID, domainTransaction.SenderID); err != nil {
			return err
		}
	} else if domainTransaction.Status == int(enums.TransactionStatusSuccess) && tempStatus == int(enums.TransactionStatusPending) {
		//Cập nhật điểm tốt của user
		var (
			user         user.User
			postType     int64
			goodDeedType int
		)

		postType, err = uc.postRepo.GetPostType(ctx, updateTransaction.InterestID)
		if err != nil {
			return err
		}

		user.ID = updateTransaction.SenderID

		if postType == int64(enums.PostTypeGiveAwayOldItem) || postType == int64(enums.PostTypeWantOldItem) {
			user.GoodPoint = enums.GoodPointGiveOldItem

			goodDeedType = int(enums.GoodDeedTypeGiveOldItem)
		} else if postType == int64(enums.PostTypeFoundItem) || postType == int64(enums.PostTypeSeekLoseItem) {
			user.GoodPoint = enums.GoodPointGiveLoseItem

			goodDeedType = int(enums.GoodDeedTypeGiveLoseItem)
		} else if postType == int64(enums.PostTypeCampaign) {
			user.GoodPoint = enums.GoodPointJoinCampaign

			goodDeedType = int(enums.GoodDeedTypeCampaign)
		}

		if err := uc.userRepo.Update(ctx, &user); err != nil {
			return err
		}

		//Tạo việc tốt cho user
		goodDeed := usergooddeed.UserGoodDeed{
			UserID:        updateTransaction.SenderID,
			GoodDeedType:  goodDeedType,
			GoodPoint:     user.GoodPoint,
			TransactionID: updateTransaction.ID,
			CreatedAt:     time.Now(),
		}

		if err := uc.userGoodDeedRepo.CreateGoodDeed(ctx, &goodDeed); err != nil {
			return err
		}
	}

	return nil
}
