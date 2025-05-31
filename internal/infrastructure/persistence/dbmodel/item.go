package dbmodel

import (
	"final_project/internal/domain/item"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"unique;size:255"`
	Description string `gorm:"type:TEXT"`
	Image       string `gorm:"type:LONGTEXT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// 1-n: Một item có nhiều item warehouse
	ItemWarehouses []ItemWarehouse `gorm:"foreignKey:ItemID"`
	// 1-n: Một item có nhiều item_import_invoice
	ItemImportInvoices []ItemImportInvoice `gorm:"foreignKey:ItemID"`
}

// Domain → DB
func ItemDomainToDB(a item.Item) Item {
	return Item{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Image:       a.Image,
	}
}

// DB → Domain
func ItemDBToDomain(a Item) item.Item {
	return item.Item{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Image:       a.Image,
	}
}
