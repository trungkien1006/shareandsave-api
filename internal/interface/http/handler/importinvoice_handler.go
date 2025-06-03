package handler

import "final_project/internal/application/importinvoiceapp"

type ImportInvoiceHandler struct {
	uc *importinvoiceapp.UseCase
}

func NewImportInvoiceHandler(uc *importinvoiceapp.UseCase) *ImportInvoiceHandler {
	return &ImportInvoiceHandler{uc: uc}
}
