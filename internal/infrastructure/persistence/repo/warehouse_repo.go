package persistence

import "gorm.io/gorm"

type WarehouseRepoDB struct {
	db *gorm.DB
}

func NewWarehouseRepoDB(db *gorm.DB) *WarehouseRepoDB {
	return &WarehouseRepoDB{db: db}
}
