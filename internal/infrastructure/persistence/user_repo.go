package persistence

import (
	"context"
	"final-project/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepoDB struct {
	db *gorm.DB
}

func NewUserRepoDB(db *gorm.DB) *UserRepoDB {
	return &UserRepoDB{db: db}
}

func (r *UserRepoDB) GetAll(ctx context.Context, users *[]user.User) error {
	if err := r.db.Debug().WithContext(ctx).Find(&users).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepoDB) GetByID(ctx context.Context, user *user.User, user_id int) error {
	if err := r.db.Debug().WithContext(ctx).Where("id = ?", user_id).First(&user).Error; err != nil {
		return err
	}

	return nil
}

// func (r *UserRepoDB) Save(ctx context.Context, u *user.User) error {
// 	_, err := r.db.ExecContext(ctx, `INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)`,
// 		u.ID, u.Name, u.Email, u.Password)
// 	return err
// }

// func (r *UserRepoDB) IsEmailExist(ctx context.Context, email string) (bool, error) {
// 	var exists bool
// 	err := r.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`, email).Scan(&exists)
// 	return exists, err
// }
