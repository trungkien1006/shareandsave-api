package requestapp

import (
	"context"
	"final_project/internal/domain/request"
	"final_project/internal/domain/user"
)

type UseCase struct {
	repo     request.Repository
	userRepo user.Repository
}

func NewUseCase(r request.Repository, userRepo user.Repository) *UseCase {
	return &UseCase{
		repo:     r,
		userRepo: userRepo,
	}
}

func (uc *UseCase) CreateRequest(ctx context.Context, req *request.Request, user *user.User) error {
	if err := uc.repo.Create(ctx, req); err != nil {
		return err
	}

	return nil
}
