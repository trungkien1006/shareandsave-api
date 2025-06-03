package persistence

import (
	"context"
	"errors"
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/domain/warehouse"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type ImportInvoiceRepoDB struct {
	db *gorm.DB
}

func NewImportInvoiceRepoDB(db *gorm.DB) *ImportInvoiceRepoDB {
	return &ImportInvoiceRepoDB{db: db}
}

func (r *ImportInvoiceRepoDB) GetImportInvoiceNum(ctx context.Context) (int, error) {
	var invoiceNum int64 = 0

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.ImportInvoice{}).
		Count(&invoiceNum).Error; err != nil {
		return int(invoiceNum + 1), errors.New("Gặp lỗi khi đếm số phiếu nhập: " + err.Error())
	}

	return int(invoiceNum + 1), nil
}

func (r *ImportInvoiceRepoDB) CreateImportInvoice(ctx context.Context, importInvoice importinvoice.ImportInvoice, warehouse *[]warehouse.Warehouse) error {
	var (
		DBImportInvoice dbmodel.ImportInvoice
		DBWarehouse     []dbmodel.Warehouse
	)

	DBImportInvoice = dbmodel.ImportInvoiceDomainToDB(importInvoice)

	for _, value := range *warehouse {
		DBWarehouse = append(DBWarehouse, dbmodel.WarehouseDomainToDB(value))
	}

	tx := r.db.Begin()

	if err := tx.Debug().WithContext(ctx).
		Model(&dbmodel.ImportInvoice{}).
		Create(&DBImportInvoice).Error; err != nil {
		return errors.New("Có lỗi khi thêm mới phiếu nhập: " + err.Error())
	}

	if err := tx.Debug().WithContext(ctx).
		Model(&dbmodel.Warehouse{}).
		Create(&DBWarehouse).Error; err != nil {
		return errors.New("Có lỗi khi thêm mới danh sách lô đồ: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("Có lỗi khi commit transaction: " + err.Error())
	}

	var tempItemName []string

	for _, value := range *warehouse {
		tempItemName = append(tempItemName, value.ItemName)
	}

	for key, value := range DBWarehouse {
		*warehouse = append(*warehouse, dbmodel.WarehouseDBToDomain(value, tempItemName[key]))
	}

	return nil
}
