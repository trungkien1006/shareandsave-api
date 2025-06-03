package importinvoice

import "context"

type Repository interface {
	GetImportInvoiceNum(ctx context.Context) (string, error)
}
