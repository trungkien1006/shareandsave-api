package interestapp

import "final_project/internal/domain/interest"

type UseCase struct {
	repo interest.Repository
}

func NewUseCase(r interest.Repository) *UseCase {
	return &UseCase{repo: r}
}
