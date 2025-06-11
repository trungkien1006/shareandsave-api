package warehouseapp

import "final_project/internal/domain/warehouse"

type UseCase struct {
	repo warehouse.Repository
}

func NewUseCase(r warehouse.Repository) *UseCase {
	return &UseCase{repo: r}
}
