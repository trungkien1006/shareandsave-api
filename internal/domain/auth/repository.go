package auth

import (
	"context"
	"final_project/internal/domain/user"
)

type Repository interface {
	SignUp(ctx context.Context, user *user.User) error
	Login(ctx context.Context, user *user.User, email, password string, isAdmin bool, clientRoleID uint) error
}
