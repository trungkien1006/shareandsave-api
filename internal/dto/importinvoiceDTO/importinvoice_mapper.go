package importinvoicedto

import (
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/pkg/enums"
)

// DTO -> Domain
func CreateDTOToDomain(dto CreateImportInvoiceRequest) importinvoice.ImportInvoice {
	return importinvoice.ImportInvoice{
		SenderID:    dto.SenderID,
		ReceiverID:  dto.ReceiverID,
		Classify:    enums.ItemClassify.String(dto.Classify),
		SendDate:    dto.SendDate,
		Description: dto.Description,
	}
}
