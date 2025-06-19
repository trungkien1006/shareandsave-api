package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/warehouse"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"final_project/internal/pkg/enums"

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
			sender.full_name AS sender_name,
			receiver.full_name AS receiver_name
		`).
		Joins("JOIN item ON item.id = warehouse.item_id").
		Joins("JOIN import_invoice as ii ON ii.id = warehouse.import_invoice_id").
		Joins("JOIN user as sender ON sender.id = ii.sender_id").
		Joins("JOIN user as receiver ON receiver.id = ii.receiver_id")

	if req.SearchBy != "" && req.SearchValue != "" {
		if req.SearchBy == "SKU" {
			req.SearchBy = "sku"
		}

		column := strcase.ToSnake(req.SearchBy) // "fullName" -> "full_name"

		if column == "sender_name" {
			column = "sender.full_name"
		} else if column == "item_name" {
			column = "item.name"
		} else if column == "receiver_name" {
			column = "receiver.receiver_name"
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
		Preload("Item").
		Preload("Item.Category")

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

func (r *WarehouseRepoDB) GetAllOldStockItem(ctx context.Context, items *[]warehouse.ItemOldStock, req filter.FilterRequest) (int, error) {
	var (
		query        *gorm.DB
		totalRecords int64
		itemOldStock []dbmodel.ItemOldStock
	)

	query = r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.ItemOldStock{}).
		Table("warehouse as w").
		Select("item.id as item_id, item.name as item_name, item.image as item_image, item.description as description, category.name as category_name, COUNT(iw.id) as quantity").
		Joins("JOIN item ON item.id = w.item_id").
		Joins("JOIN category ON category.id = item.category_id").
		Joins("JOIN item_warehouse as iw ON iw.warehouse_id = w.id").
		Where("w.classify = ?", enums.ItemClassifyOlItem).
		Group("item.id, item.name, item.image, item.description, category.name")

	if req.SearchBy != "" && req.SearchValue != "" {
		column := strcase.ToSnake(req.SearchBy) // "fullName" -> "full_name"

		if column == "item_name" {
			column = "item.name"
		} else if column == "description" {
			column = "item.description"
		} else if column == "category_name" {
			column = "category.name"
		} else {
			column = "w." + column
		}

		query.Where(column+" LIKE ? ", "%"+req.SearchValue+"%")
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, err
	}

	if req.Sort != "" && req.Order != "" {
		query = query.Order(strcase.ToSnake(req.Sort) + " " + req.Order)
	}

	if req.Limit > 0 && req.Page > 0 {
		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}

	if err := query.Find(&itemOldStock).Error; err != nil {
		return 0, err
	}

	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

	// mappingOb := make(map[uint]warehouse.ItemOldStock, 0)

	// for _, value := range itemOldStock {
	// 	mappingOb[value.ItemID] = append(*items, dbmodel.ItemOldStockDBToDomain(value))
	// }

	for _, value := range itemOldStock {
		*items = append(*items, dbmodel.ItemOldStockDBToDomain(value))
	}

	return totalPages, nil
}

func (r *WarehouseRepoDB) GetItemByCode(ctx context.Context, itemWarehouse *warehouse.ItemWareHouse, code string) error {
	var dbItemWarehouse dbmodel.ItemWarehouse

	if err := r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.ItemWarehouse{}).
		Table("item_warehouse as iw").
		Where("iw.code = ?", code).
		Joins("JOIN item ON item.id = iw.item_id").
		Preload("Item").
		Preload("Item.Category").
		Find(&dbItemWarehouse).Error; err != nil {
		return errors.New("Có lỗi khi lấy đồ theo mã code: " + err.Error())
	}

	*itemWarehouse = dbmodel.ItemWarehouseDBToDomain(dbItemWarehouse)

	return nil
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
			sender.full_name AS sender_name,
			receiver.full_name AS receiver_name
		`).
		Preload("ItemWarehouses").
		Preload("ItemWarehouses.Item").
		Preload("ItemWarehouses.Item.Category").
		Where("warehouse.id = ?", warehouseID).
		Joins("JOIN item_warehouse as iw ON iw.warehouse_id = warehouse.id").
		Joins("JOIN item ON item.id = warehouse.item_id").
		Joins("JOIN import_invoice as ii ON ii.id = warehouse.import_invoice_id").
		Joins("JOIN user as sender ON sender.id = ii.sender_id").
		Joins("JOIN user as receiver ON receiver.id = ii.receiver_id").
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
