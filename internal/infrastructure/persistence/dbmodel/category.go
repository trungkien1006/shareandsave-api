package dbmodel

import (
	"final_project/internal/domain/category"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Items []Item `gorm:"foreignKey:CategoryID"`
}

// DB -> Domain
func DBToDomain(db Category) category.Category {
	return category.Category{
		ID:   db.ID,
		Name: db.Name,
	}
}

// Domain -> DB
func DomainToDB(domain category.Category) Category {
	return Category{
		ID:   domain.ID,
		Name: domain.Name,
	}
}
