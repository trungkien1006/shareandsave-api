package dbmodel

import (
	"final_project/internal/domain/item"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	CategoryID  uint   `gorm:"index"` // mới thêm
	Name        string `gorm:"unique;size:255"`
	Description string `gorm:"type:TEXT"`
	Image       string `gorm:"type:LONGTEXT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Category           Category            `gorm:"foreignKey:CategoryID"`
	ItemWarehouses     []ItemWarehouse     `gorm:"foreignKey:ItemID"`
	ItemImportInvoices []ItemImportInvoice `gorm:"foreignKey:ItemID"`
	TransactionItems   []TransactionItem   `gorm:"foreignKey:ItemID"`
}

// Domain → DB
func ItemDomainToDB(a item.Item) Item {
	return Item{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Image:       a.Image,
		CategoryID:  a.CategoryID,
	}
}

// DB → Domain
func ItemDBToDomain(a Item) item.Item {
	return item.Item{
		ID:           a.ID,
		Name:         a.Name,
		Description:  a.Description,
		Image:        a.Image,
		CategoryID:   a.CategoryID,
		CategoryName: a.Category.Name,
	}
}
