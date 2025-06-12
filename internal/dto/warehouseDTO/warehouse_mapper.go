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

// Domain to DTO
func DetailWarehouseDomainToDTO(domain warehouse.DetailWarehouse) DetailWarehouseDTO {
	itemWarehouses := make([]ItemWareHouseDTO, 0)

	for _, value := range domain.ItemWareHouse {
		itemWarehouses = append(itemWarehouses, ItemWareHouseDTO{
			ID:          value.ID,
			ItemID:      value.ItemID,
			ItemName:    value.ItemName,
			WarehouseID: value.WarehouseID,
			Code:        value.Code,
			Description: value.Description,
			Status:      value.Status,
		})
	}

	return DetailWarehouseDTO{
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
		ItemWareHouse:   itemWarehouses,
	}
}
