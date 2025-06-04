package importinvoicedto

import (
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/pkg/enums"
)

// DTO -> Domain
func CreateDTOToDomain(dto CreateImportInvoiceRequest) importinvoice.ImportInvoice {
	var items []importinvoice.ItemImportInvoice

	for _, value := range dto.ItemImportInvoice {
		items = append(items, importinvoice.ItemImportInvoice{
			ItemID:      value.ItemID,
			Quantity:    value.Quantity,
			Description: value.Description,
		})
	}

	return importinvoice.ImportInvoice{
		SenderID:          dto.SenderID,
		Classify:          enums.ItemClassify.String(dto.Classify),
		Description:       dto.Description,
		ItemImportInvoice: items,
	}
}

// Domain -> DTO
func GetDomainToDTO(domain importinvoice.GetImportInvoice) ImportInvoiceListDTO {
	return ImportInvoiceListDTO{
		ID:           domain.ID,
		SenderName:   domain.SenderName,
		ReceiverName: domain.ReceiverName,
		Classify:     domain.Classify,
		CreatedAt:    domain.CreatedAt,
		ItemCount:    domain.ItemCount,
	}
}
