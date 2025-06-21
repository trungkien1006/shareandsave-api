package importinvoicedto

import (
	importinvoice "final_project/internal/domain/import_invoice"
	warehousedto "final_project/internal/dto/warehouseDTO"
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
		Classify:          int(dto.Classify),
		Description:       dto.Description,
		ItemImportInvoice: items,
	}
}

// Domain -> DTO
func GetDomainToDTO(domain importinvoice.GetImportInvoice) ImportInvoiceListDTO {
	return ImportInvoiceListDTO{
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
func ImportInvoiceDomainToDTO(domain importinvoice.ImportInvoice) ImportInvoiceDTO {
	warehouses := make([]warehousedto.DetailWarehouseDTO, 0)
	items := make([]ItemImportInvoiceDTO, 0)

	for _, value := range domain.ItemImportInvoice {
		items = append(items, ItemImportInvoiceDTO{
			ItemID:      value.ItemID,
			ItemName:    value.ItemName,
			Quantity:    value.Quantity,
			Description: value.Description,
		})
	}

	for _, v := range domain.Warehouses {
		var itemWarehouses []warehousedto.ItemWareHouseDTO

		for _, value := range v.ItemWareHouse {
			itemWarehouses = append(itemWarehouses, warehousedto.ItemWareHouseDTO{
				ItemID:      value.ItemID,
				ItemName:    value.ItemName,
				Code:        value.Code,
				Description: value.Description,
				Status:      value.Status,
			})
		}

		warehouses = append(warehouses, warehousedto.DetailWarehouseDTO{
			ItemID:        v.ItemID,
			SKU:           v.SKU,
			Quantity:      v.Quantity,
			Classify:      v.Classify,
			Description:   v.Description,
			StockPlace:    v.StockPlace,
			ItemWareHouse: itemWarehouses,
		})
	}

	return ImportInvoiceDTO{
		InvoiceNum:        domain.InvoiceNum,
		SenderID:          domain.SenderID,
		ReceiverID:        domain.ReceiverID,
		Classify:          domain.Classify,
		Description:       domain.Description,
		IsLock:            domain.IsLock,
		ItemImportInvoice: items,
		Warehouses:        warehouses,
	}
}
