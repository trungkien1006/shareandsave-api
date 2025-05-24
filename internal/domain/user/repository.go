package user

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, users *[]User, req filter.FilterRequest) (int, error)
	GetByID(ctx context.Context, user *User, user_id int) error
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	IsEmailExist(ctx context.Context, email string) (bool, error)
	IsPhoneNumberExist(ctx context.Context, phoneNumber string) (bool, error)
}
