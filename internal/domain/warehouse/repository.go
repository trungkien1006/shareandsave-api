package warehouse

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, warehouses *[]Warehouse, req filter.FilterRequest) (int, error)
	GetAllItem(ctx context.Context, warehouses *[]ItemWareHouse, req filter.FilterRequest) (int, error)
	GetAllOldStockItem(ctx context.Context, items *[]ItemOldStock, req filter.FilterRequest) (int, error)
	GetItemByCode(ctx context.Context, itemWarehouse *ItemWareHouse, code string) error
	GetByID(ctx context.Context, warehouse *DetailWarehouse, warehouseID uint) error
	GetItemWarehouseQuantity(ctx context.Context, itemID uint) (uint, error)
	Update(ctx context.Context, warehouse DetailWarehouse) error
	IsExist(ctx context.Context, itemWarehouseID uint) (bool, error)
	GetSKUByItemWarehouseID(ctx context.Context, itemWarehouseID uint) (string, error)
}
