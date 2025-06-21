package exportinvoicedto

import exportinvoice "final_project/internal/domain/export_invoice"

// Domain -> DTO
func GetDomainToDTO(domain exportinvoice.GetExportInvoice) ExportInvoiceListDTO {
	return ExportInvoiceListDTO{
		ID:           domain.ID,
		InvoiceNum:   domain.InvoiceNum,
		SenderName:   domain.SenderName,
		ReceiverName: domain.ReceiverName,
		Classify:     domain.Classify,
		CreatedAt:    domain.CreatedAt,
		ItemCount:    domain.ItemCount,
	}
}
