package auth

import (
	"context"
	"final_project/internal/domain/user"
)

type Repository interface {
	Login(ctx context.Context, user *user.User, email, password string) ([]string, error)
}
