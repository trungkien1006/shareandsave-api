package categoryapp

import (
	"context"
	"final_project/internal/domain/category"
)

type UseCase struct {
	repo category.Repository
}

func NewUseCase(r category.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllCategories(ctx context.Context, categories *[]category.Category) error {
	if err := uc.repo.GetAllCategories(ctx, categories); err != nil {
		return err
	}

	return nil
}
