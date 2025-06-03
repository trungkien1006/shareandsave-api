package importinvoice

import (
	"context"
	"final_project/internal/domain/warehouse"
)

type Repository interface {
	GetImportInvoiceNum(ctx context.Context) (int, error)
	CreateImportInvoice(ctx context.Context, importInvoice ImportInvoice, warehouse *[]warehouse.Warehouse) error
}
