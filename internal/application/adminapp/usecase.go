package adminapp

import (
	"context"
	"errors"
	"final_project/internal/domain/admin"
	"final_project/internal/domain/filter"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
)

type UseCase struct {
	repo admin.Repository
}

func NewUseCase(r admin.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllAdmin(ctx context.Context, admins *[]admin.Admin, domainReq filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, admins, domainReq)
	if err != nil {
		return 0, err
	}
	return totalPage, nil
}

func (uc *UseCase) GetAdminByID(ctx context.Context, adminObj *admin.Admin, adminID int) error {
	if err := uc.repo.GetByID(ctx, adminObj, adminID); err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) CreateAdmin(ctx context.Context, adminObj *admin.Admin) error {
	roleExisted, err := uc.repo.IsRoleExist(ctx, adminObj.RoleID)
	if err != nil {
		return err
	}
	if !roleExisted {
		return errors.New(enums.ErrRoleNotExist)
	}

	emailExisted, err := uc.repo.IsEmailExist(ctx, adminObj.Email)
	if err != nil {
		return err
	}
	if emailExisted {
		return errors.New(enums.ErrEmailExisted)
	}

	hashedPassword, err := hash.HashPassword(adminObj.Password)
	if err != nil {
		return err
	}
	adminObj.Password = hashedPassword

	if err := uc.repo.Save(ctx, adminObj); err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) UpdateAdmin(ctx context.Context, domainAdmin *admin.Admin) error {
	var updateAdmin admin.Admin

	if err := uc.repo.GetByID(ctx, &updateAdmin, int(domainAdmin.ID)); err != nil {
		return err
	}

	if updateAdmin.ID == 0 {
		return errors.New(enums.ErrUserNotExist)
	}

	if domainAdmin.FullName != "" {
		updateAdmin.FullName = domainAdmin.FullName
	}

	updateAdmin.Status = domainAdmin.Status
	updateAdmin.ID = domainAdmin.ID

	if domainAdmin.RoleID != 0 {
		roleExisted, err := uc.repo.IsRoleExist(ctx, domainAdmin.RoleID)
		if err != nil {
			return err
		}
		if !roleExisted {
			return errors.New(enums.ErrRoleNotExist)
		}

		updateAdmin.RoleID = domainAdmin.RoleID
	}
	if domainAdmin.Password != "" {
		hashedPassword, err := hash.HashPassword(domainAdmin.Password)
		if err != nil {
			return err
		}
		updateAdmin.Password = hashedPassword
	}

	if err := uc.repo.Update(ctx, &updateAdmin); err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) DeleteAdmin(ctx context.Context, adminID int) error {
	var deleteAdmin admin.Admin

	if err := uc.repo.GetByID(ctx, &deleteAdmin, int(adminID)); err != nil {
		return err
	}

	if deleteAdmin.ID == 0 {
		return errors.New(enums.ErrUserNotExist)
	}

	if err := uc.repo.Delete(ctx, &deleteAdmin); err != nil {
		return err
	}
	return nil
}
