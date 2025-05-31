package user

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, users *[]User, req filter.FilterRequest, roleID uint) (int, error)
	GetByID(ctx context.Context, domainUser *User, userID int, roleID uint) error
	GetIDByEmailPhoneNumber(ctx context.Context, email string, phoneNumber string) (uint, error)
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, domainUser *User) error
	IsEmailExist(ctx context.Context, email string, userID int) (bool, error)
	IsPhoneNumberExist(ctx context.Context, phoneNumber string, userID int) (bool, error)
	IsTableEmpty(ctx context.Context) (bool, error)
}
