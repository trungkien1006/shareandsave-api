package persistence

import (
	"context"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/warehouse"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type WarehouseRepoDB struct {
	db *gorm.DB
}

func NewWarehouseRepoDB(db *gorm.DB) *WarehouseRepoDB {
	return &WarehouseRepoDB{db: db}
}

func (r *WarehouseRepoDB) GetAll(ctx context.Context, warehouses *[]warehouse.Warehouse, req filter.FilterRequest) (int, error) {
	var (
		query        *gorm.DB
		totalRecords int64
		dbWarehouse  []dbmodel.DetailWarehouse
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.DetailWarehouse{}).
		Table("warehouse").
		Select(`
			warehouse.*,
			item.name AS item_name,
			sender.full_name AS sender_name
		`).
		Joins("JOIN item_warehouse as iw ON iw.warehouse_id = warehouse.id").
		Joins("JOIN item ON item.id = warehouse.item_id").
		Joins("JOIN import_invoice as ii ON ii.id = warehouse.import_invoice_id").
		Joins("JOIN user as sender ON sender.id = ii.sender_id")

	if req.SearchBy != "" && req.SearchValue != "" {
		column := strcase.ToSnake(req.SearchBy) // "fullName" -> "full_name"

		if column == "sender_name" {
			column = "sender.full_name"
		} else if column == "item_name" {
			column = "item.name"
		} else {
			column = "warehouse." + column
		}

		query.Where(column+" LIKE ? ", "%"+req.SearchValue+"%")
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, err
	}

	if req.Sort != "" && req.Order != "" {
		query = query.Order("warehouse." + strcase.ToSnake(req.Sort) + " " + req.Order)
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&dbWarehouse).Error; err != nil {
		return 0, err
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	for _, value := range dbWarehouse {
		*warehouses = append(*warehouses, dbmodel.DetailDBToDomain(value))
	}

	return totalPages, nil
}
