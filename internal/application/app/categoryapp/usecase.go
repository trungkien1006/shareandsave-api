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

func (uc *UseCase) CreateCategory(ctx context.Context, category *category.Category) error {
	if err := uc.repo.Save(ctx, category); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdateCategory(ctx context.Context, category *category.Category) error {
	if err := uc.repo.Update(ctx, category); err != nil {
		return err
	}

	return nil
}
