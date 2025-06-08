package transactionapp

import (
	"context"
	"final_project/internal/domain/transaction"
)

type UseCase struct {
	repo transaction.Repository
}

func NewUseCase(r transaction.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) CreateTransaction(ctx context.Context, transaction *transaction.Transaction) error {
	return nil
}
