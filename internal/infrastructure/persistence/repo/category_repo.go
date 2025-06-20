package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/category"
	"final_project/internal/domain/item"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type CategoryRepoDB struct {
	db *gorm.DB
}

func NewCategoryRepoDB(db *gorm.DB) *CategoryRepoDB {
	return &CategoryRepoDB{db: db}
}

func (r *CategoryRepoDB) Save(ctx context.Context, category *category.Category) error {
	dbCategory := dbmodel.DomainToDB(*category)

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Category{}).Create(&dbCategory).Error; err != nil {
		return errors.New("Có lỗi khi tạo loại đồ đạc: " + err.Error())
	}

	return nil
}

func (r *CategoryRepoDB) GetAllCategories(ctx context.Context, categories *[]category.Category) error {
	var dbCategories []dbmodel.Category

	if err := r.db.Debug().WithContext(ctx).Find(&dbCategories).Error; err != nil {
		return err
	}

	for _, value := range dbCategories {
		*categories = append(*categories, dbmodel.DBToDomain(value))
	}

	return nil
}

func (r *CategoryRepoDB) GetCategoryNameByItemIDs(ctx context.Context, itemIDsMap map[uint]uint, categoryName *[]string) error {
	var (
		itemIDs []uint
	)

	for _, value := range itemIDsMap {
		itemIDs = append(itemIDs, value)
	}

	if err := r.db.Debug().WithContext(ctx).
		Model(&item.Item{}).
		Table("item as item").
		Joins("JOIN category ON category.id = item.category_id").
		Where("item.id IN ?", itemIDs).
		Distinct("category.name").
		Pluck("category.name", &categoryName).Error; err != nil {
		return errors.New("Lỗi khi truy vấn tên category theo danh sách item ID: " + err.Error())
	}

	return nil
}

func (r *CategoryRepoDB) IsTableEmpty(ctx context.Context) (bool, error) {
	var count int64

	if err := r.db.WithContext(ctx).Model(&category.Category{}).Count(&count).Error; err != nil {
		return false, err
	}

	return count == 0, nil
}
