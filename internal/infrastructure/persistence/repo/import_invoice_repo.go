package persistence

import (
	"context"
	"errors"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"strconv"

	"gorm.io/gorm"
)

type ImportInvoiceRepoDB struct {
	db *gorm.DB
}

func NewImportInvoiceRepoDB(db *gorm.DB) *ImportInvoiceRepoDB {
	return &ImportInvoiceRepoDB{db: db}
}

func (r *ImportInvoiceRepoDB) GetImportInvoiceNum(ctx context.Context) (string, error) {
	var invoiceNum int64 = 0

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.ImportInvoice{}).
		Count(&invoiceNum).Error; err != nil {
		return "PN" + strconv.Itoa(int(invoiceNum+1)), errors.New("Gặp lỗi khi đếm số phiếu nhập: " + err.Error())
	}

	return "PN" + strconv.Itoa(int(invoiceNum+1)), nil
}
