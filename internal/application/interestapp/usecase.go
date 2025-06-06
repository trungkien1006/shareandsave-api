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
