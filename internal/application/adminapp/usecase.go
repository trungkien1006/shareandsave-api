package adminapp

import (
	"context"
	"errors"
	"final_project/internal/domain/admin"
	"final_project/internal/domain/filter"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
)

type UseCase struct {
	repo     admin.Repository
	roleRepo rolepermission.Repository
}

func NewUseCase(r admin.Repository, roleRepo rolepermission.Repository) *UseCase {
	return &UseCase{
		repo:     r,
		roleRepo: roleRepo,
	}
}

func (uc *UseCase) GetAllAdmin(ctx context.Context, filter filter.FilterRequest) ([]struct {
	Admin    admin.Admin
	RoleName string
}, int, error) {
	dbAdminWithRoles, totalPage, err := uc.repo.GetAllWithRole(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var result []struct {
		Admin    admin.Admin
		RoleName string
	}

	for _, dbAdminWithRole := range dbAdminWithRoles {
		result = append(result, struct {
			Admin    admin.Admin
			RoleName string
		}{
			Admin:    dbAdminWithRole.Admin,
			RoleName: dbAdminWithRole.RoleName,
		})
	}
	return result, totalPage, nil
}

func (uc *UseCase) GetAdminByID(ctx context.Context, adminID uint) (admin.Admin, string, error) {
	dbAdminWithRole, err := uc.repo.GetByIDWithRole(ctx, adminID)

	if err != nil {
		return admin.Admin{}, "", err
	}

	return dbAdminWithRole.Admin, dbAdminWithRole.RoleName, nil
}

func (uc *UseCase) CreateAdmin(ctx context.Context, domainAdmin admin.Admin) (admin.Admin, string, error) {
	roleExisted, err := uc.repo.IsRoleExist(ctx, domainAdmin.RoleID)
	if err != nil {
		return admin.Admin{}, "", err
	}
	if !roleExisted {
		return admin.Admin{}, "", errors.New(enums.ErrRoleNotExist)
	}

	emailExisted, err := uc.repo.IsEmailExist(ctx, domainAdmin.Email)
	if err != nil {
		return admin.Admin{}, "", err
	}
	if emailExisted {
		return admin.Admin{}, "", errors.New(enums.ErrEmailExisted)
	}

	hashedPassword, err := hash.HashPassword(domainAdmin.Password)
	if err != nil {
		return admin.Admin{}, "", err
	}
	domainAdmin.Password = hashedPassword

	// Gọi repo để lưu
	createdDomainAdmin, err := uc.repo.Save(ctx, domainAdmin)

	if err != nil {
		return admin.Admin{}, "", err
	}

	// Lấy lại tên role (nếu cần)
	roleName, err := uc.roleRepo.GetRoleNameByID(ctx, createdDomainAdmin.RoleID)

	if err != nil {
		return admin.Admin{}, "", err
	}

	return createdDomainAdmin, roleName, nil
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
