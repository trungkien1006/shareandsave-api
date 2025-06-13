package warehouse

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, warehouses *[]Warehouse, req filter.FilterRequest) (int, error)
	GetAllItem(ctx context.Context, warehouses *[]ItemWareHouse, req filter.FilterRequest) (int, error)
	GetByID(ctx context.Context, warehouse *DetailWarehouse, warehouseID uint) error
	Update(ctx context.Context, warehouse DetailWarehouse) error
}
