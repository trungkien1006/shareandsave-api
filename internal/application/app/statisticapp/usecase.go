package statisticapp

import (
	"context"
	"final_project/internal/domain/statistic"
)

type UseCase struct {
	repo statistic.Repository
}

func NewUseCase(repo statistic.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) TotalTransaction(ctx context.Context) (int64, int64, error) {
	total, totalLastMonth, err := uc.repo.TotalTransaction(ctx)
	if err != nil {
		return 0, 0, err
	}

	return total, totalLastMonth, nil
}
