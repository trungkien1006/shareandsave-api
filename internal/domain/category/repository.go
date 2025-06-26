package category

import "context"

type Repository interface {
	Save(ctx context.Context, category *Category) error
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, categoryID uint) error
	IsTableEmpty(ctx context.Context) (bool, error)
	GetAllCategories(ctx context.Context, categories *[]Category) error
	GetCategoryNameByItemIDs(ctx context.Context, itemIDs map[uint]uint, categoryName *[]string) error
}
