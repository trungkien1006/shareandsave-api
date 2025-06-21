package exportinvoicedto

import (
	exportinvoice "final_project/internal/domain/export_invoice"
)

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

// Domain to DTO
func ExportInvoiceDomainToDTO(domain exportinvoice.ExportInvoice) ExportInvoiceDTO {
	items := make([]ItemExportInvoiceDTO, 0)

	for _, value := range domain.ItemExportInvoices {
		items = append(items, ItemExportInvoiceDTO{
			ID:              value.ID,
			InvoiceID:       value.InvoiceID,
			ItemWarehouseID: value.ItemWarehouseID,
			SKU:             value.SKU,
		})
	}

	return ExportInvoiceDTO{
		ID:                domain.ID,
		InvoiceNum:        domain.InvoiceNum,
		SenderID:          domain.SenderID,
		SenderName:        domain.SenderName,
		ReceiverID:        domain.ReceiverID,
		ReceiverName:      domain.ReceiverName,
		Classify:          domain.Classify,
		Description:       domain.Description,
		IsLock:            domain.IsLock,
		ItemExportInvoice: items,
	}
}

// DTO to Domain
func ExportInvoiceDTOToDomain(dto CreateExportInvoiceRequest) exportinvoice.ExportInvoice {
	items := make([]exportinvoice.ItemExportInvoice, 0)

	for _, value := range dto.ItemExportInvoice {
		items = append(items, exportinvoice.ItemExportInvoice{
			ItemWarehouseID: value.ItemWarehouseID,
		})
	}

	return exportinvoice.ExportInvoice{
		Classify:           int(dto.Classify),
		SenderID:           dto.SenderID,
		Description:        dto.Description,
		ItemExportInvoices: items,
	}
}
