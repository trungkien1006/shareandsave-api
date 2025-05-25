package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/item"
	"final_project/internal/pkg/helpers"
	"final_project/internal/reference"

	"gorm.io/gorm"
)

type ItemRepoDB struct {
	db *gorm.DB
}

func NewItemRepoDB(db *gorm.DB) *ItemRepoDB {
	return &ItemRepoDB{db: db}
}

func (r *ItemRepoDB) Save(ctx context.Context, i *item.Item) error {
	return r.db.Debug().WithContext(ctx).Create(i).Error
}

func (r *ItemRepoDB) GetAll(ctx context.Context, items *[]item.Item, req filter.FilterRequest) (int, error) {
	var (
		tableName    = "item"
		query        *gorm.DB
		totalRecords int64
	)

	query = r.db.Debug().WithContext(ctx).Model(&item.Item{})

	if req.Filter != "" {
		var filters []reference.FilterStruc
		err := json.Unmarshal([]byte(req.Filter), &filters)
		if err != nil {
			return 0, errors.New("Lỗi khi chuyển đổi filter từ JSON thành struct: " + err.Error())
		}
		helpers.Filter(query, filters, tableName)
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, err
	}

	if req.Sort != "" && req.Order != "" {
		query = query.Order(req.Sort + " " + req.Order)
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(items).Error; err != nil {
		return 0, err
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	return totalPages, nil
}

func (r *ItemRepoDB) GetByID(ctx context.Context, item *item.Item, id uint) error {
	if err := r.db.Debug().WithContext(ctx).First(&item, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *ItemRepoDB) Update(ctx context.Context, i *item.Item) error {
	return r.db.Debug().WithContext(ctx).Save(i).Error
}

func (r *ItemRepoDB) Delete(ctx context.Context, i *item.Item) error {
	return r.db.Debug().WithContext(ctx).Delete(i).Error
}
