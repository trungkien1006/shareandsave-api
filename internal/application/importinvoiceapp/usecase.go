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
