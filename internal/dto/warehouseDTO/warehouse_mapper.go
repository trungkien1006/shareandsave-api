package warehousedto

import "final_project/internal/domain/warehouse"

// Domain to DTO
func WarehouseDomainToDTO(domain warehouse.Warehouse) WarehouseDTO {
	return WarehouseDTO{
		ID:              domain.ID,
		ItemID:          domain.ItemID,
		ImportInvoiceID: domain.ImportInvoiceID,
		SenderName:      domain.SenderName,
		ItemName:        domain.ItemName,
		SKU:             domain.SKU,
		Quantity:        domain.Quantity,
		Description:     domain.Description,
		Classify:        domain.Classify,
		StockPlace:      domain.StockPlace,
	}
}
