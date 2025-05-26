package admin

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, admins *[]Admin, req filter.FilterRequest) (int, error)
	GetByID(ctx context.Context, domainAdmin *Admin, adminID int) error
	Save(ctx context.Context, admin *Admin) error
	Update(ctx context.Context, admin *Admin) error
	Delete(ctx context.Context, domainAdmin *Admin) error
	IsEmailExist(ctx context.Context, email string) (bool, error)
}
