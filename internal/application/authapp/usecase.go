package authapp

import (
	"context"
	"final_project/internal/domain/auth"
	"final_project/internal/domain/user"
)

type UseCase struct {
	repo auth.Repository
}

func NewUseCase(r auth.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) Login(ctx context.Context, domainAuthLogin *auth.AuthLogin, domainUser *user.User) error {

	return nil
}
