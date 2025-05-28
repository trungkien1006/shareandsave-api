package admin

import (
	"context"
	"final_project/internal/domain/filter"
)

type AdminWithRole struct {
	Admin    Admin
	RoleName string
}

type Repository interface {
	GetAllWithRole(ctx context.Context, filter filter.FilterRequest) ([]AdminWithRole, int, error)
	GetAll(ctx context.Context, admins *[]Admin, req filter.FilterRequest) (int, error)
	GetByID(ctx context.Context, domainAdmin *Admin, adminID int) error
	GetByIDWithRole(ctx context.Context, adminID uint) (AdminWithRole, error)
	Save(ctx context.Context, domainAdmin Admin) (Admin, error)
	Update(ctx context.Context, admin *Admin) error
	Delete(ctx context.Context, domainAdmin *Admin) error
	IsEmailExist(ctx context.Context, email string) (bool, error)
	IsRoleExist(ctx context.Context, roleID uint) (bool, error)
	IsTableEmpty(ctx context.Context) (bool, error)
}
