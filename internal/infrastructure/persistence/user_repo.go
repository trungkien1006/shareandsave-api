package persistence

import (
	"context"
	"encoding/json"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/helpers"
	"final_project/internal/reference"
	"math"

	"gorm.io/gorm"
)

type UserRepoDB struct {
	db *gorm.DB
}

func NewUserRepoDB(db *gorm.DB) *UserRepoDB {
	return &UserRepoDB{db: db}
}

func (r *UserRepoDB) GetAll(ctx context.Context, users *[]user.User, req filter.FilterRequest) (int, error) {
	var tableName = "user"
	var query *gorm.DB

	query = r.db.Debug().WithContext(ctx).Table(tableName)

	if req.Filter != "" {
		var filters []reference.FilterStruc

		err := json.Unmarshal([]byte(req.Filter), &filters)
		if err != nil {
			return 0, err
		}

		helpers.Filter(query, filters, tableName)
	}

	var totalRecord int64 = 0

	//lay ra tong so record
	if err := query.Count(&totalRecord).Error; err != nil {
		return 0, err
	}

	//phan trang
	if req.Limit != 0 && req.Page != 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	//sort du lieu
	if req.Sort != "" {
		query.Order(req.Sort + " " + req.Order)
	}

	//tinh toan total page
	totalPage := int(math.Ceil(float64(totalRecord) / float64(req.Limit)))

	return totalPage, nil
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
