package warehouseapp

import (
	"context"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/warehouse"
)

type UseCase struct {
	repo warehouse.Repository
}

func NewUseCase(r warehouse.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllWarehouse(ctx context.Context, warehouses *[]warehouse.Warehouse, filter filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, warehouses, filter)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetAllItemWarehouse(ctx context.Context, warehouses *[]warehouse.ItemWareHouse, filter filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAllItem(ctx, warehouses, filter)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetItemByCode(ctx context.Context, itemWarehouse *warehouse.ItemWareHouse, code string) error {
	if err := uc.repo.GetItemByCode(ctx, itemWarehouse, code); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetWarehouseByID(ctx context.Context, warehouse *warehouse.DetailWarehouse, warehouseID uint) error {
	if err := uc.repo.GetByID(ctx, warehouse, warehouseID); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdateWarehouse(ctx context.Context, domainWarehouse warehouse.DetailWarehouse) error {
	var updateWarehouse warehouse.DetailWarehouse

	if domainWarehouse.Description != "" {
		updateWarehouse.Description = domainWarehouse.Description
	}

	if domainWarehouse.StockPlace != "" {
		updateWarehouse.StockPlace = domainWarehouse.StockPlace
	}

	if domainWarehouse.ItemWareHouse != nil {
		updateWarehouse.ItemWareHouse = domainWarehouse.ItemWareHouse
	}

	if err := uc.repo.Update(ctx, updateWarehouse); err != nil {
		return err
	}

	return nil
}
