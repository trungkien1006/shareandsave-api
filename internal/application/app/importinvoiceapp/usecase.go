package importinvoiceapp

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/filter"
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/domain/item"
	"final_project/internal/domain/redis"
	"final_project/internal/domain/user"
	"final_project/internal/domain/warehouse"
	"final_project/internal/pkg/enums"
	"strconv"
)

type UseCase struct {
	repo      importinvoice.Repository
	service   *importinvoice.ImportInvoiceService
	userRepo  user.Repository
	itemRepo  item.Repository
	redisRepo redis.Repository
}

func NewUseCase(r importinvoice.Repository, userRepo user.Repository, itemRepo item.Repository, service *importinvoice.ImportInvoiceService, redisRepo redis.Repository) *UseCase {
	return &UseCase{
		repo:      r,
		userRepo:  userRepo,
		itemRepo:  itemRepo,
		service:   service,
		redisRepo: redisRepo,
	}
}

func (uc *UseCase) GetAllImportInvoice(ctx context.Context, importInvoice *[]importinvoice.GetImportInvoice, filter filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, importInvoice, filter)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) CreateImportInvoice(ctx context.Context, importInvoice *importinvoice.ImportInvoice) error {
	var (
		handlerWarehouse []warehouse.DetailWarehouse
	)

	// Lấy số hóa đơn hiện tại
	invoiceNum, err := uc.repo.GetImportInvoiceNum(ctx)
	if err != nil {
		return err
	}

	importInvoice.InvoiceNum = invoiceNum
	importInvoice.IsLock = false

	// Kiểm tra người gửi có tồn tại hay không
	senderExisted, err := uc.userRepo.IsExist(ctx, importInvoice.SenderID)
	if err != nil {
		return err
	}

	if !senderExisted {
		return errors.New("Người gửi không tồn tại")
	}

	// Kiểm tra món đồ có tồn tại hay không
	for key, value := range importInvoice.ItemImportInvoice {
		var item item.Item

		err := uc.itemRepo.GetByID(ctx, &item, value.ItemID)
		if err != nil {
			return err
		}

		if item.ID == 0 {
			return errors.New("Món đồ không tồn tại")
		}

		importInvoice.ItemImportInvoice[key].ItemName = item.Name
		importInvoice.ItemImportInvoice[key].MaxClaim = item.MaxClaim
	}

	// Gom nhóm các món đồ thành 1 lô và tạo danh sách các món đồ thuộc lô
	warehouses := make(map[uint]warehouse.DetailWarehouse)

	for _, value := range importInvoice.ItemImportInvoice {
		if wh, ok := warehouses[value.ItemID]; ok {
			wh.Quantity = wh.Quantity + int(value.Quantity)

			warehouses[value.ItemID] = wh
		} else {
			wh := warehouses[value.ItemID]

			wh.ItemID = value.ItemID
			wh.ItemName = value.ItemName
			wh.MaxClaim = value.MaxClaim
			wh.SKU = uc.service.GenerateSKU(int(value.ItemID))
			wh.Classify = importInvoice.Classify
			wh.Description = ""
			wh.Quantity = int(value.Quantity)
			wh.StockPlace = ""

			warehouses[value.ItemID] = wh
		}

		var itemWHs []warehouse.ItemWareHouse

		for i := 0; i < int(value.Quantity); i++ {
			itemCode, err := uc.service.GenerateUniqueDigitString(9)
			if err != nil {
				return errors.New("Có lỗi khi tạo mã code cho món đồ: " + err.Error())
			}

			itemWHs = append(itemWHs, warehouse.ItemWareHouse{
				ItemID:      value.ItemID,
				ItemName:    value.ItemName,
				MaxClaim:    value.MaxClaim,
				Description: value.Description,
				Code:        itemCode,
				Status:      int(enums.ItemWarehouseStatusInStock),
			})
		}

		wh := warehouses[value.ItemID]

		wh.ItemWareHouse = itemWHs

		warehouses[value.ItemID] = wh
	}

	for _, value := range warehouses {
		handlerWarehouse = append(handlerWarehouse, value)
	}

	importInvoice.Warehouses = handlerWarehouse

	if err := uc.repo.CreateImportInvoice(ctx, importInvoice); err != nil {
		return err
	}

	//Update quantity item in redis
	for _, value := range importInvoice.ItemImportInvoice {
		itemClaimJSON, err := uc.redisRepo.GetFromRedisHash(ctx, enums.ItemClaimRequest, "item:"+strconv.Itoa(int(value.ItemID)))
		if err != nil {
			return errors.New("Có lỗi khi thực hiện tăng số lượng đồ trong redis hash map: " + err.Error())
		}

		var itemClaim warehouse.ClaimRequestItem

		if itemClaimJSON == "" {
			err = json.Unmarshal([]byte(itemClaimJSON), &itemClaim)
			if err != nil {
				return errors.New("Có lỗi khi thực hiện tăng số lượng đồ trong redis hash map, decode JSON error: " + err.Error())
			}
		}

		itemClaim.ItemQuantity += uint(value.Quantity)
		itemClaim.MaxClaim = uint(value.MaxClaim)

		newItemClaimJSON, err := json.Marshal(itemClaim)
		if err != nil {
			return errors.New("Có lỗi khi thực hiện tăng số lượng đồ trong redis hash map, encode JSON error: " + err.Error())
		}

		if err := uc.redisRepo.SetToRedisHash(ctx, enums.ItemClaimRequest, "item:"+strconv.Itoa(int(value.ItemID)), string(newItemClaimJSON)); err != nil {
			return err
		}
	}

	return nil
}
