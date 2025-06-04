package importinvoice

import (
	"context"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/warehouse"
)

type Repository interface {
	GetAll(ctx context.Context, importInvoice *[]GetImportInvoice, filter filter.FilterRequest) (int, error)
	GetImportInvoiceNum(ctx context.Context) (int, error)
	CreateImportInvoice(ctx context.Context, importInvoice ImportInvoice, warehouse *[]warehouse.Warehouse) error
}
