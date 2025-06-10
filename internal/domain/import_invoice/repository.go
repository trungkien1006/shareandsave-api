package importinvoice

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, importInvoice *[]GetImportInvoice, filter filter.FilterRequest) (int, error)
	GetImportInvoiceNum(ctx context.Context) (int, error)
	CreateImportInvoice(ctx context.Context, importInvoice *ImportInvoice) error
	IsTableEmpty(ctx context.Context) (bool, error)
}
