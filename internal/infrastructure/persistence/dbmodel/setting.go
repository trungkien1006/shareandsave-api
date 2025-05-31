package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Setting struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Key       string `gorm:"size:255"`
	Value     string `gorm:"type:MEDIUMTEXT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
