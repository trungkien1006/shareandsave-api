package importinvoiceapp

import (
	"context"
	"errors"
	"final_project/internal/domain/filter"
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/domain/item"
	"final_project/internal/domain/user"
	"final_project/internal/domain/warehouse"
	"final_project/internal/pkg/enums"
)

type UseCase struct {
	repo     importinvoice.Repository
	service  *importinvoice.ImportInvoiceService
	userRepo user.Repository
	itemRepo item.Repository
}

func NewUseCase(r importinvoice.Repository, userRepo user.Repository, itemRepo item.Repository, service *importinvoice.ImportInvoiceService) *UseCase {
	return &UseCase{
		repo:     r,
		userRepo: userRepo,
		itemRepo: itemRepo,
		service:  service,
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
	}

	// Gom nhóm các món đồ thành 1 lô và tạo danh sách các món đồ thuộc lô
	warehouses := make(map[uint]warehouse.DetailWarehouse)

	for _, value := range importInvoice.ItemImportInvoice {
		if wh, ok := warehouses[value.ID]; ok {
			wh.Quantity = wh.Quantity + int(value.Quantity)

			warehouses[value.ID] = wh
		} else {
			wh := warehouses[value.ID]

			wh.ItemID = value.ItemID
			wh.ItemName = value.ItemName
			wh.SKU = uc.service.GenerateSKU(int(value.ID))
			wh.Classify = importInvoice.Classify
			wh.Description = ""
			wh.Quantity = int(value.Quantity)
			wh.StockPlace = ""

			warehouses[value.ID] = wh
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
				Description: value.Description,
				Code:        itemCode,
				Status:      int(enums.ItemWarehouseStatusInStock),
			})
		}

		wh := warehouses[value.ID]

		wh.ItemWareHouse = itemWHs

		warehouses[value.ID] = wh
	}

	for _, value := range warehouses {
		handlerWarehouse = append(handlerWarehouse, value)
	}

	importInvoice.Warehouses = handlerWarehouse

	if err := uc.repo.CreateImportInvoice(ctx, importInvoice); err != nil {
		return err
	}

	return nil
}
