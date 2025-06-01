package item

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	GetAll(ctx context.Context, items *[]Item, req filter.FilterRequest) (int, error)
	GetByID(ctx context.Context, item *Item, id uint) error
	IsExisted(ctx context.Context, itemID uint) (bool, error)
	Save(ctx context.Context, item *Item) error
	Update(ctx context.Context, item *Item) error
	Delete(ctx context.Context, item *Item) error
	IsTableEmpty(ctx context.Context) (bool, error)
}
