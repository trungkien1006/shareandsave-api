package exportinvoiceapp

import exportinvoice "final_project/internal/domain/export_invoice"

type UseCase struct {
	repo exportinvoice.Repository
}

func NewUseCase(r exportinvoice.Repository) *UseCase {
	return &UseCase{
		repo: r,
	}
}
