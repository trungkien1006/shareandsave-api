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
