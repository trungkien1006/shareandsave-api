package authapp

import (
	"final_project/internal/domain/auth"
)

type UseCase struct {
	repo auth.Repository
}

func NewUseCase(r auth.Repository) *UseCase {
	return &UseCase{repo: r}
}
