package warehousedto

import "final_project/internal/domain/warehouse"

// Domain to DTO
func ItemWarehouseDomainToDTO(domain warehouse.ItemWareHouse) ItemWarehouse {
	return ItemWarehouse{
		ItemID:      domain.ItemID,
		ItemName:    domain.ItemName,
		Code:        domain.Code,
		Description: domain.Description,
	}
}
