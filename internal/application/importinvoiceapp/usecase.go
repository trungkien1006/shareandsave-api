package importinvoiceapp

import (
	importinvoice "final_project/internal/domain/import_invoice"
)

type UseCase struct {
	repo importinvoice.Repository
}

func NewUseCase(r importinvoice.Repository) *UseCase {
	return &UseCase{repo: r}
}

// func (uc *UseCase) CreateImportInvoice(ctx context.Context, importInvoice importinvoice.ImportInvoice) error {
// 	// Lấy số hóa đơn hiện tại
// 	invoiceNum, err := uc.repo.GetImportInvoiceNum(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
