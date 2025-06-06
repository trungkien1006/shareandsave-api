package interestapp

import (
	"context"
	"final_project/internal/domain/interest"
)

type UseCase struct {
	repo interest.Repository
}

func NewUseCase(r interest.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllInterest(ctx context.Context, postInterest *[]interest.PostInterest, userID uint, filter interest.GetInterest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, postInterest, userID, filter)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) CreateInterest(ctx context.Context, interest interest.Interest) error {
	if err := uc.repo.Create(ctx, interest); err != nil {
		return err
	}

	return nil
}
