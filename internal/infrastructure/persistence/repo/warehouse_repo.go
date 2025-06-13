package persistence

import (
	"context"
	"errors"
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

func (r *WarehouseRepoDB) GetAllItem(ctx context.Context, itemWarehouses *[]warehouse.ItemWareHouse, req filter.FilterRequest) (int, error) {
	var (
		query           *gorm.DB
		totalRecords    int64
		dbItemWarehouse []dbmodel.ItemWarehouse
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.ItemWarehouse{}).
		Table("item_warehouse as iw").
		Joins("JOIN item ON item.id = iw.item_id").
		Preload("Item")

	if req.SearchBy != "" && req.SearchValue != "" {
		column := strcase.ToSnake(req.SearchBy) // "fullName" -> "full_name"

		if column == "item_name" {
			column = "item.name"
		} else {
			column = "iw." + column
		}

		query.Where(column+" LIKE ? ", "%"+req.SearchValue+"%")
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, err
	}

	if req.Sort != "" && req.Order != "" {
		query = query.Order("iw." + strcase.ToSnake(req.Sort) + " " + req.Order)
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&dbItemWarehouse).Error; err != nil {
		return 0, err
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	for _, value := range dbItemWarehouse {
		*itemWarehouses = append(*itemWarehouses, dbmodel.ItemWarehouseDBToDomain(value))
	}

	return totalPages, nil
}

func (r *WarehouseRepoDB) GetByID(ctx context.Context, warehouse *warehouse.DetailWarehouse, warehouseID uint) error {
	var dbWarehouse dbmodel.DetailWarehouse

	if err := r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.DetailWarehouse{}).
		Table("warehouse").
		Select(`
			warehouse.*,
			item.name AS item_name,
			sender.full_name AS sender_name
		`).
		Where("warehouse.id = ?", warehouseID).
		Joins("JOIN item_warehouse as iw ON iw.warehouse_id = warehouse.id").
		Joins("JOIN item ON item.id = warehouse.item_id").
		Joins("JOIN import_invoice as ii ON ii.id = warehouse.import_invoice_id").
		Joins("JOIN user as sender ON sender.id = ii.sender_id").
		First(&dbWarehouse).Error; err != nil {
		return errors.New("Có lỗi khi truy vấn warehouse: " + err.Error())
	}

	*warehouse = dbmodel.DetailDBToDetailDomain(dbWarehouse)

	return nil
}

func (r *WarehouseRepoDB) Update(ctx context.Context, warehouse warehouse.DetailWarehouse) error {
	var dbWarehouse dbmodel.Warehouse

	tx := r.db.Debug().WithContext(ctx).Begin()

	dbWarehouse = dbmodel.UpdateDomainToDB(warehouse)

	if err := tx.Model(&dbmodel.Warehouse{}).
		Where("id = ?", dbWarehouse.ID).
		Updates(&dbWarehouse).Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi cập nhật lô hàng: " + err.Error())
	}

	for _, value := range dbWarehouse.ItemWarehouses {
		if err := tx.Model(&dbmodel.ItemWarehouse{}).
			Where("id = ?", value.ID).
			Updates(&value).Error; err != nil {
			tx.Rollback()
			return errors.New("Có lỗi khi cập nhật chi tiết lô hàng: " + err.Error())
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi commit transaction: " + err.Error())
	}

	return nil
}
