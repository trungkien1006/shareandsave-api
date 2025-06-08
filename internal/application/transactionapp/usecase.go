package transactionapp

import (
	"final_project/internal/domain/transaction"
)

type UseCase struct {
	repo transaction.Repository
}

func NewUseCase(r transaction.Repository) *UseCase {
	return &UseCase{repo: r}
}
