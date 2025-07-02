package transactionapp

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/domain/item"
	"final_project/internal/domain/notification"
	"final_project/internal/domain/post"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/transaction"
	"final_project/internal/domain/user"
	usergooddeed "final_project/internal/domain/user_good_deed"
	"final_project/internal/pkg/enums"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type UseCase struct {
	repo             transaction.Repository
	userRepo         user.Repository
	interestRepo     interest.Repository
	itemRepo         item.Repository
	postRepo         post.Repository
	userGoodDeedRepo usergooddeed.Repository
	roleRepo         rolepermission.Repository
	clientID         uint
	notiService      notification.Service
}

func NewUseCase(r transaction.Repository, userRepo user.Repository, interestRepo interest.Repository, itemRepo item.Repository, postRepo post.Repository, userGoodDeedRepo usergooddeed.Repository, roleRepo rolepermission.Repository, notiService notification.Service) *UseCase {
	ctx := context.Background()

	clientID, err := roleRepo.GetRoleIDByName(ctx, "Client")
	if err != nil {
		fmt.Println("Có lỗi khi set clientID cho user usecase: " + err.Error())
	}

	return &UseCase{
		repo:             r,
		userRepo:         userRepo,
		interestRepo:     interestRepo,
		itemRepo:         itemRepo,
		postRepo:         postRepo,
		userGoodDeedRepo: userGoodDeedRepo,
		roleRepo:         roleRepo,
		clientID:         clientID,
		notiService:      notiService,
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

	// noti := notification.Notification{
	// 	SenderID: transaction.,
	// }

	return nil
}

func (uc *UseCase) UpdateTransaction(ctx context.Context, domainTransaction *transaction.Transaction) error {
	var updateTransaction transaction.Transaction

	notiContent := "Giao dịch: "
	isChange := false

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
		if tempStatus == int(enums.TransactionStatusPending) && domainTransaction.Status == int(enums.TransactionStatusRollBack) {
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
		isChange = true
	}

	if domainTransaction.Items != nil {
		updateTransaction.Items = domainTransaction.Items

		notiContent += "cập nhật đồ đạc, "

		isChange = true
	}

	if domainTransaction.Method != "" {
		updateTransaction.Method = domainTransaction.Method

		notiContent += "cập nhật phương thức trao đổi, "

		isChange = true
	}

	// Cập nhật
	if err := uc.repo.Update(ctx, &updateTransaction); err != nil {
		return err
	}

	if domainTransaction.Status == int(enums.TransactionStatusRollBack) && tempStatus == int(enums.TransactionStatusSuccess) {
		var (
			postType   int64
			updateUser user.User
		)

		postType, err = uc.postRepo.GetPostType(ctx, updateTransaction.InterestID)
		if err != nil {
			return err
		}

		if err := uc.userRepo.GetUserByID(ctx, &updateUser, int(updateTransaction.SenderID), uc.clientID, 0); err != nil {
			return err
		}

		if updateUser.ID == 0 {
			return errors.New(enums.ErrUserNotExist)
		}

		if postType == int64(enums.PostTypeGiveAwayOldItem) {
			updateUser.GoodPoint -= enums.GoodPointGiveOldItem
		} else if postType == int64(enums.PostTypeFoundItem) {
			updateUser.GoodPoint -= enums.GoodPointGiveLoseItem
		} else if postType == int64(enums.PostTypeCampaign) {
			updateUser.GoodPoint -= enums.GoodPointJoinCampaign
		}

		if err := uc.userRepo.Update(ctx, &updateUser); err != nil {
			return err
		}

		//Xóa việc tốt cho user
		if err := uc.userGoodDeedRepo.DeleteGoodDeed(ctx, domainTransaction.ID, domainTransaction.SenderID); err != nil {
			return err
		}
	} else if domainTransaction.Status == int(enums.TransactionStatusSuccess) && tempStatus == int(enums.TransactionStatusPending) {
		//Cập nhật điểm tốt của user
		var (
			updateUser   user.User
			postType     int64
			goodDeedType int
		)

		postType, err = uc.postRepo.GetPostType(ctx, updateTransaction.InterestID)
		if err != nil {
			return err
		}

		if err := uc.userRepo.GetUserByID(ctx, &updateUser, int(updateTransaction.SenderID), uc.clientID, 0); err != nil {
			return err
		}

		if updateUser.ID == 0 {
			return errors.New(enums.ErrUserNotExist)
		}

		if postType == int64(enums.PostTypeGiveAwayOldItem) || postType == int64(enums.PostTypeWantOldItem) {
			updateUser.GoodPoint += enums.GoodPointGiveOldItem

			goodDeedType = int(enums.GoodDeedTypeGiveOldItem)
		} else if postType == int64(enums.PostTypeFoundItem) || postType == int64(enums.PostTypeSeekLoseItem) {
			updateUser.GoodPoint += enums.GoodPointGiveLoseItem

			goodDeedType = int(enums.GoodDeedTypeGiveLoseItem)
		} else if postType == int64(enums.PostTypeCampaign) {
			updateUser.GoodPoint += enums.GoodPointJoinCampaign

			goodDeedType = int(enums.GoodDeedTypeCampaign)
		}

		if err := uc.userRepo.Update(ctx, &updateUser); err != nil {
			return err
		}

		//Tạo việc tốt cho user
		goodDeed := usergooddeed.UserGoodDeed{
			UserID:        updateTransaction.SenderID,
			GoodDeedType:  goodDeedType,
			GoodPoint:     updateUser.GoodPoint,
			TransactionID: &updateTransaction.ID,
			CreatedAt:     time.Now(),
		}

		if err := uc.userGoodDeedRepo.CreateGoodDeed(ctx, &goodDeed); err != nil {
			return err
		}
	}

	if isChange {
		noti := notification.Notification{
			Type:       "normal",
			TargetType: "transaction",
			TargetID:   updateTransaction.ID,
			IsRead:     false,
			SenderID:   &updateTransaction.SenderID,
			ReceiverID: &updateTransaction.ReceiverID,
		}

		if tempStatus != updateTransaction.Status {
			switch updateTransaction.Status {
			case int(enums.TransactionStatusCancelled):
				notiContent += "đã bị hủy, "
				break
			case int(enums.TransactionStatusSuccess):
				notiContent += "đã được xác nhận, "
				break
			case int(enums.TransactionStatusRollBack):
				notiContent += "đã bị hoàn tác, "
				break
			}
		}

		noti.Content = strings.TrimSuffix(notiContent, ",") + "!"

		if err := uc.notiService.CreateAndPushSocket(ctx, &noti); err != nil {
			return errors.New("Có lỗi khi thêm thông báo: " + err.Error())
		}
	}

	return nil
}
