package persistence

import (
	"context"
	"errors"
	exportinvoice "final_project/internal/domain/export_invoice"
	"final_project/internal/domain/filter"
	"math"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type ExportInvoiceRepoDB struct {
	db *gorm.DB
}

func NewExportInvoiceRepoDB(db *gorm.DB) *ExportInvoiceRepoDB {
	return &ExportInvoiceRepoDB{db: db}
}

func (r *ExportInvoiceRepoDB) GetAll(ctx context.Context, exportInvoice *[]exportinvoice.GetExportInvoice, filter filter.FilterRequest) (uint, error) {
	var (
		query *gorm.DB
	)

	query = r.db.Debug().WithContext(ctx).
		Table("export_invoice as ei").
		Select(`
		ei.id,
		ei.invoice_num,
		sender.full_name AS sender_name,
		receiver.full_name AS receiver_name,
		SUM(iei.quantity) AS item_count,
		ei.created_at,
		ei.classify
	`).
		Joins("LEFT JOIN user AS sender ON sender.id = ei.sender_id").
		Joins("LEFT JOIN user AS receiver ON receiver.id = ei.receiver_id").
		Joins("LEFT JOIN item_export_invoice AS iei ON iei.invoice_id = ei.id").
		Group("ei.id, sender.full_name, receiver.full_name, ei.created_at, ei.classify")

	if filter.SearchBy != "" && filter.SearchValue != "" {
		column := strcase.ToSnake(filter.SearchBy) // "fullName" -> "full_name"

		if column == "sender_name" {
			column = "sender.full_name"
		} else if column == "receiver_name" {
			column = "receiver.full_name"
		} else {
			column = "ei." + column
		}

		query.Where(column+" LIKE ? ", "%"+filter.SearchValue+"%")

	}

	var totalRecord int64 = 0

	//lay ra tong so record
	if err := query.Count(&totalRecord).Error; err != nil {
		return 0, errors.New("Lỗi khi đếm tổng số record của phiếu xuất kho: " + err.Error())
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
			filter.Sort = "ei." + strcase.ToSnake(filter.Sort)
		}

		query.Order(filter.Sort + " " + filter.Order)
	}

	if err := query.Find(&exportInvoice).Error; err != nil {
		return 0, errors.New("Lỗi khi lọc phiếu xuất kho: " + err.Error())
	}

	//tinh toan total page
	totalPage := int(math.Ceil(float64(totalRecord) / float64(filter.Limit)))

	return uint(totalPage), nil
}
