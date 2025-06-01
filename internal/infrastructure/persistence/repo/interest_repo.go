package persistence

import "gorm.io/gorm"

type InterestRepoDB struct {
	db *gorm.DB
}

func NewInterestRepoDB(db *gorm.DB) *InterestRepoDB {
	return &InterestRepoDB{db: db}
}
