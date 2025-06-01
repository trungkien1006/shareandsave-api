package categoryapp

import (
	"final_project/internal/domain/category"
)

type UseCase struct {
	repo category.Repository
}

func NewUseCase(r category.Repository) *UseCase {
	return &UseCase{repo: r}
}
