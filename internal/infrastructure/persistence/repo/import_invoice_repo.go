package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/filter"
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"math"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type ImportInvoiceRepoDB struct {
	db *gorm.DB
}

func NewImportInvoiceRepoDB(db *gorm.DB) *ImportInvoiceRepoDB {
	return &ImportInvoiceRepoDB{db: db}
}

func (r *ImportInvoiceRepoDB) GetAll(ctx context.Context, importInvoice *[]importinvoice.GetImportInvoice, filter filter.FilterRequest) (int, error) {
	var (
		query *gorm.DB
	)

	query = r.db.Debug().WithContext(ctx).
		Table("import_invoice as ii").
		Select(`
		ii.id,
		sender.full_name AS sender_name,
		receiver.full_name AS receiver_name,
		SUM(iii.quantity) AS item_count,
		ii.created_at,
		ii.classify
	`).
		Joins("LEFT JOIN user AS sender ON sender.id = ii.sender_id").
		Joins("LEFT JOIN user AS receiver ON receiver.id = ii.receiver_id").
		Joins("LEFT JOIN item_import_invoice AS iii ON iii.invoice_id = ii.id").
		Group("ii.id, sender.full_name, receiver.full_name, ii.created_at, ii.classify")

	if filter.SearchBy != "" && filter.SearchValue != "" {
		column := strcase.ToSnake(filter.SearchBy) // "fullName" -> "full_name"

		if column == "sender_name" {
			column = "sender.full_name"
		} else if column == "receiver_name" {
			column = "receiver.full_name"
		} else {
			column = "ii." + column
		}

		query.Where(column+" LIKE ? ", "%"+filter.SearchValue+"%")

	}

	var totalRecord int64 = 0

	//lay ra tong so record
	if err := query.Count(&totalRecord).Error; err != nil {
		return 0, errors.New("Lỗi khi đếm tổng số record của bài viết: " + err.Error())
	}

	//phan trang
	if filter.Limit != 0 && filter.Page != 0 {
		query.Offset((filter.Page - 1) * filter.Limit).Limit(filter.Limit)
	}

	//sort du lieu
	if filter.Sort != "" {
		if filter.Sort == "itemCount" {
			filter.Sort = "item_count"
		} else {
			filter.Sort = "ii." + strcase.ToSnake(filter.Sort)
		}

		query.Order(filter.Sort + " " + filter.Order)
	}

	if err := query.Find(&importInvoice).Error; err != nil {
		return 0, errors.New("Lỗi khi lọc phiếu nhập kho: " + err.Error())
	}

	//tinh toan total page
	totalPage := int(math.Ceil(float64(totalRecord) / float64(filter.Limit)))

	return totalPage, nil
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

func (r *ImportInvoiceRepoDB) CreateImportInvoice(ctx context.Context, importInvoice *importinvoice.ImportInvoice) error {
	var (
		dbImportInvoice dbmodel.ImportInvoice
	)

	dbImportInvoice = dbmodel.ImportInvoiceDomainToDB(*importInvoice)

	tx := r.db.Begin()

	if err := tx.Debug().WithContext(ctx).
		Model(&dbmodel.ImportInvoice{}).
		Create(&dbImportInvoice).Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi thêm mới phiếu nhập: " + err.Error())
	}

	// if err := tx.Debug().WithContext(ctx).
	// 	Model(&dbmodel.Warehouse{}).
	// 	Create(&DBWarehouse).Error; err != nil {
	// 	tx.Rollback()
	// 	return errors.New("Có lỗi khi thêm mới danh sách lô đồ: " + err.Error())
	// }

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("Có lỗi khi commit transaction: " + err.Error())
	}

	*importInvoice = dbmodel.ImportInvoiceDBToDomain(dbImportInvoice)

	return nil
}

func (r *ImportInvoiceRepoDB) IsTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&dbmodel.ImportInvoice{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
