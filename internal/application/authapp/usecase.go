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

func (uc *UseCase) Login(ctx context.Context, domainAuthLogin auth.AuthLogin, JWT *string, refreshToken *string, domainUser *user.User) error {
	if err := uc.repo.Login(ctx, domainUser, domainAuthLogin.Email, domainAuthLogin.Password); err != nil {
		return err
	}

	return nil
}
