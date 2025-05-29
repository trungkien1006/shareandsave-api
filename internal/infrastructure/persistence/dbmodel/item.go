package dbmodel

import (
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
	// Relations
	Posts          []Post          `gorm:"foreignKey:ItemID"`
	ItemWarehouses []ItemWarehouse `gorm:"foreignKey:ItemID"`
}

// Domain → DB
func ItemDomainToDB(a Item) Item {
	return Item{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Image:       a.Image,
	}
}

// DB → Domain
func ItemDBToDomain(a Item) Item {
	return Item{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Image:       a.Image,
	}
}
