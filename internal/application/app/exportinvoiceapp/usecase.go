package exportinvoiceapp

import (
	"context"
	exportinvoice "final_project/internal/domain/export_invoice"
	"final_project/internal/domain/filter"
)

type UseCase struct {
	repo exportinvoice.Repository
}

func NewUseCase(r exportinvoice.Repository) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) GetAllExportInvoice(ctx context.Context, exportInvoice *[]exportinvoice.GetExportInvoice, filter filter.FilterRequest) (uint, error) {
	totalPage, err := uc.repo.GetAll(ctx, exportInvoice, filter)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}
