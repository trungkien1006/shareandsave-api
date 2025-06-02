package authapp

import (
	"context"
	"final_project/internal/domain/auth"
)

type UseCase struct {
	repo auth.Repository
}

func NewUseCase(r auth.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) Login(ctx context.Context, domainAuthLogin auth.AuthLogin, JWT *string, refreshToken *string) error {

	return nil
}
