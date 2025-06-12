package warehouse

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, warehouses *[]Warehouse, req filter.FilterRequest) (int, error)
}
