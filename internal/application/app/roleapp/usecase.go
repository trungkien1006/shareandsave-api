package roleapp

import (
	"context"
	rolepermission "final_project/internal/domain/role_permission"
)

type UseCase struct {
	repo rolepermission.Repository
}

func NewUseCase(r rolepermission.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllRoles(ctx context.Context, roles *[]rolepermission.Role) error {
	if err := uc.repo.GetAllRoles(ctx, roles); err != nil {
		return err
	}

	return nil
}
