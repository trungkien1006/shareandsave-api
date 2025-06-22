package exportinvoiceapp

import (
	"context"
	"encoding/json"
	"errors"
	exportinvoice "final_project/internal/domain/export_invoice"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/redis"
	"final_project/internal/domain/user"
	"final_project/internal/domain/warehouse"
	"final_project/internal/pkg/enums"
	"strconv"
)

type UseCase struct {
	repo      exportinvoice.Repository
	userRepo  user.Repository
	whRepo    warehouse.Repository
	redisRepo redis.Repository
}

func NewUseCase(r exportinvoice.Repository, userRepo user.Repository, whRepo warehouse.Repository, redisRepo redis.Repository) *UseCase {
	return &UseCase{
		repo:      r,
		userRepo:  userRepo,
		whRepo:    whRepo,
		redisRepo: redisRepo,
	}
}

func (uc *UseCase) GetAllExportInvoice(ctx context.Context, exportInvoice *[]exportinvoice.GetExportInvoice, filter filter.FilterRequest) (uint, error) {
	totalPage, err := uc.repo.GetAll(ctx, exportInvoice, filter)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) Create(ctx context.Context, exportInvoice *exportinvoice.ExportInvoice) error {
	// Lấy số hóa đơn hiện tại
	invoiceNum, err := uc.repo.GetExportInvoiceNum(ctx)
	if err != nil {
		return err
	}

	exportInvoice.InvoiceNum = invoiceNum
	exportInvoice.IsLock = false

	// Kiểm tra người gửi có tồn tại hay không
	senderExisted, err := uc.userRepo.IsExist(ctx, exportInvoice.SenderID)
	if err != nil {
		return err
	}

	if !senderExisted {
		return errors.New("Người gửi không tồn tại")
	}

	// Kiểm tra món đồ có tồn tại hay không và truy xuất sku
	for key, value := range exportInvoice.ItemExportInvoices {
		sku, err := uc.whRepo.GetSKUByItemWarehouseID(ctx, value.ItemWarehouseID)
		if err != nil {
			return err
		}

		if sku == "" {
			return errors.New("SKU của lô hàng rỗng")
		}

		exportInvoice.ItemExportInvoices[key].SKU = sku
	}

	if err := uc.repo.Create(ctx, exportInvoice); err != nil {
		return err
	}

	//Update quantity item in redis
	for _, value := range exportInvoice.ItemExportInvoices {
		var itemID int = 0

		if err := uc.whRepo.GetItemIDByItemWarehouseID(ctx, value.ItemWarehouseID, &itemID); err != nil {
			return err
		}

		itemClaimJSON, err := uc.redisRepo.GetFromRedisHash(ctx, enums.ItemClaimRequest, "item:"+strconv.Itoa(itemID))
		if err != nil {
			return errors.New("Có lỗi khi thực hiện tăng số lượng đồ trong redis hash map: " + err.Error())
		}

		var itemClaim warehouse.ClaimRequestItem

		err = json.Unmarshal([]byte(itemClaimJSON), &itemClaim)
		if err != nil {
			return errors.New("Có lỗi khi thực hiện tăng số lượng đồ trong redis hash map, decode JSON error: " + err.Error())
		}

		itemClaim.ItemQuantity -= 1

		newItemClaimJSON, err := json.Marshal(itemClaim)
		if err != nil {
			return errors.New("Có lỗi khi thực hiện tăng số lượng đồ trong redis hash map, encode JSON error: " + err.Error())
		}

		if err := uc.redisRepo.SetToRedisHash(ctx, enums.ItemClaimRequest, "item:"+strconv.Itoa(itemID), string(newItemClaimJSON)); err != nil {
			return err
		}
	}

	return nil
}
