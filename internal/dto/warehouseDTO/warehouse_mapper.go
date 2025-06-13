package warehousedto

import "final_project/internal/domain/warehouse"

// DTO to Domain
func UpdateWarehouseDTOToDomain(dto UpdateWarehouseRequest) warehouse.DetailWarehouse {
	var itemWarehouses []warehouse.ItemWareHouse

	for _, value := range dto.ItemWarehouses {
		itemWarehouses = append(itemWarehouses, warehouse.ItemWareHouse{
			ID:          value.ID,
			Description: value.Descripton,
		})
	}

	return warehouse.DetailWarehouse{
		Description:   dto.Description,
		StockPlace:    dto.StockPlace,
		ItemWareHouse: itemWarehouses,
	}
}

// Domain to DTO
func WarehouseDomainToDTO(domain warehouse.Warehouse) WarehouseDTO {
	return WarehouseDTO{
		ID:              domain.ID,
		ItemID:          domain.ItemID,
		ImportInvoiceID: domain.ImportInvoiceID,
		SenderName:      domain.SenderName,
		ReceiverName:    domain.ReceiverName,
		ItemName:        domain.ItemName,
		SKU:             domain.SKU,
		Quantity:        domain.Quantity,
		Description:     domain.Description,
		Classify:        domain.Classify,
		StockPlace:      domain.StockPlace,
		CreatedAt:       domain.CreatedAt,
	}
}

// Domain to DTO
func ItemWarehouseDomainToDTO(domain warehouse.ItemWareHouse) ItemWareHouseDTO {
	return ItemWareHouseDTO{
		ID:           domain.ID,
		ItemID:       domain.ItemID,
		ItemName:     domain.ItemName,
		CategoryName: domain.CategoryName,
		WarehouseID:  domain.WarehouseID,
		Code:         domain.Code,
		Description:  domain.Description,
		Status:       domain.Status,
	}
}

// Domain to DTO
func DetailWarehouseDomainToDTO(domain warehouse.DetailWarehouse) DetailWarehouseDTO {
	itemWarehouses := make([]ItemWareHouseDTO, 0)

	for _, value := range domain.ItemWareHouse {
		itemWarehouses = append(itemWarehouses, ItemWareHouseDTO{
			ID:           value.ID,
			ItemID:       value.ItemID,
			ItemName:     value.ItemName,
			CategoryName: value.CategoryName,
			WarehouseID:  value.WarehouseID,
			Code:         value.Code,
			Description:  value.Description,
			Status:       value.Status,
		})
	}

	return DetailWarehouseDTO{
		ID:              domain.ID,
		ItemID:          domain.ItemID,
		ImportInvoiceID: domain.ImportInvoiceID,
		SenderName:      domain.SenderName,
		ReceiverName:    domain.ReceiverName,
		ItemName:        domain.ItemName,
		SKU:             domain.SKU,
		Quantity:        domain.Quantity,
		Description:     domain.Description,
		Classify:        domain.Classify,
		StockPlace:      domain.StockPlace,
		ItemWareHouse:   itemWarehouses,
		CreatedAt:       domain.CreatedAt,
	}
}
