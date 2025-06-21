package exportinvoiceapp

import (
	"context"
	"errors"
	exportinvoice "final_project/internal/domain/export_invoice"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	"final_project/internal/domain/warehouse"
)

type UseCase struct {
	repo     exportinvoice.Repository
	userRepo user.Repository
	whRepo   warehouse.Repository
}

func NewUseCase(r exportinvoice.Repository, userRepo user.Repository, whRepo warehouse.Repository) *UseCase {
	return &UseCase{
		repo:     r,
		userRepo: userRepo,
		whRepo:   whRepo,
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

	// if err := uc.repo.CreateImportInvoice(ctx, exportInvoice); err != nil {
	// 	return err
	// }

	return nil
}
