package transactionapp

import (
	"context"
	"final_project/internal/domain/transaction"
	"final_project/internal/domain/user"
)

type UseCase struct {
	repo     transaction.Repository
	userRepo user.Repository
}

func NewUseCase(r transaction.Repository, userRepo user.Repository) *UseCase {
	return &UseCase{
		repo:     r,
		userRepo: userRepo,
	}
}

func (uc *UseCase) CreateTransaction(ctx context.Context, transaction *transaction.Transaction) error {
	return nil
}
