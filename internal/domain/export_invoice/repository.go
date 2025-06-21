package exportinvoice

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, exportInvoice *[]GetExportInvoice, filter filter.FilterRequest) (uint, error)
	GetExportInvoiceNum(ctx context.Context) (int, error)
}
