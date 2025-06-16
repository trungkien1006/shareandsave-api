package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/item"
	itemdto "final_project/internal/dto/itemDTO"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type ItemRepoDB struct {
	db *gorm.DB
}

func NewItemRepoDB(db *gorm.DB) *ItemRepoDB {
	return &ItemRepoDB{db: db}
}

func (r *ItemRepoDB) Save(ctx context.Context, i *item.Item) error {
	dbItem := dbmodel.ItemDomainToDB(*i)

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Item{}).Create(&dbItem).Error; err != nil {
		return errors.New("Có lỗi khi thêm item mới: " + err.Error())
	}

	*i = dbmodel.ItemDBToDomain(dbItem)

	return nil
}

func (r *ItemRepoDB) GetAll(ctx context.Context, items *[]item.Item, req filter.FilterRequest) (int, error) {
	var (
		query        *gorm.DB
		totalRecords int64
		dbItems      []dbmodel.Item
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Item{}).
		Preload("Category").
		// Table("item as item").
		Joins("JOIN category ON category.id = item.category_id")

	if req.SearchBy != "" && req.SearchValue != "" {
		column := strcase.ToSnake(req.SearchBy) // "fullName" -> "full_name"

		if column == "category_name" {
			column = "category.name"
		} else {
			column = "item." + column
		}

		query.Where(column+" LIKE ? ", "%"+req.SearchValue+"%")

	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, err
	}

	if req.Sort != "" && req.Order != "" {
		query = query.Order("item." + strcase.ToSnake(req.Sort) + " " + req.Order)
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&dbItems).Error; err != nil {
		return 0, err
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	for _, value := range dbItems {
		*items = append(*items, dbmodel.ItemDBToDomain(value))
	}

	return totalPages, nil
}

func (r *ItemRepoDB) GetByID(ctx context.Context, item *item.Item, id uint) error {
	var DBItem dbmodel.Item

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Item{}).First(&DBItem, id).Error; err != nil {
		return errors.New("Có lỗi khi truy suất đồ đạc theo id: " + err.Error())
	}

	*item = dbmodel.ItemDBToDomain(DBItem)

	return nil
}

func (r *ItemRepoDB) GetByName(ctx context.Context, item *item.Item, name string) error {
	var dbItem dbmodel.Item

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Item{}).Where("name = ?", name).First(&dbItem).Error; err != nil {
		return errors.New("Có lỗi khi truy suất đồ đạc theo tên: " + err.Error())
	}

	*item = dbmodel.ItemDBToDomain(dbItem)

	return nil
}

// func (r *ItemRepoDB) GetCategoryNameByItemID(ctx context.Context, itemID uint) (string, error) {
// 	var item

// 	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Item{}).Where("id = ?", itemID).Preload("Category").Find
// }

func (r *ItemRepoDB) IsExist(ctx context.Context, itemID uint) (bool, error) {
	var count int64

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Item{}).Where("id = ?", itemID).Count(&count).Error; err != nil {
		return false, errors.New("Có lỗi khi kiểm tra item tồn tại: " + err.Error())
	}

	return count > 0, nil
}

func (r *ItemRepoDB) Update(ctx context.Context, i *item.Item) error {
	var dbItem dbmodel.Item

	dbItem = itemdto.DomainItemToDTO(*i)

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Item{}).
		Omit("CreatedAt").
		Omit("DeleteAt").
		Where("id = ?", dbItem.ID).
		Updates(&dbItem).Error; err != nil {
		return errors.New("Có lỗi khi cập nhật item: " + err.Error())
	}

	return nil
}

func (r *ItemRepoDB) Delete(ctx context.Context, i *item.Item) error {
	return r.db.Debug().WithContext(ctx).Delete(i).Error
}

func (r *ItemRepoDB) IsTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&item.Item{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
