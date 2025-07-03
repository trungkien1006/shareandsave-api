package statisticapp

import (
	"context"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/statistic"
	"fmt"
)

type UseCase struct {
	repo     statistic.Repository
	roleRepo rolepermission.Repository
	clientID uint
}

func NewUseCase(repo statistic.Repository, roleRepo rolepermission.Repository) *UseCase {
	ctx := context.Background()

	clientID, err := roleRepo.GetRoleIDByName(ctx, "Client")
	if err != nil {
		fmt.Println("Có lỗi khi set clientID cho user usecase: " + err.Error())
	}

	return &UseCase{
		repo:     repo,
		roleRepo: roleRepo,
		clientID: clientID,
	}
}

func (uc *UseCase) TotalTransaction(ctx context.Context) (int64, int64, error) {
	total, totalLastMonth, err := uc.repo.TotalTransaction(ctx)
	if err != nil {
		return 0, 0, err
	}

	return total, totalLastMonth, nil
}

func (uc *UseCase) TotalUser(ctx context.Context) (int64, int64, error) {
	total, totalLastMonth, err := uc.repo.TotalUser(ctx, uc.clientID)
	if err != nil {
		return 0, 0, err
	}

	return total, totalLastMonth, nil
}
