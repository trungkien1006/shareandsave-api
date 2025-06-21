package dbmodel

import (
	"final_project/internal/domain/setting"
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

// Domain to DB
func SettingDomainToDB(domain setting.Setting) Setting {
	return Setting{
		ID:    domain.ID,
		Key:   domain.Key,
		Value: domain.Value,
	}
}

// DB to Domain
func SettingDBToDomain(db Setting) setting.Setting {
	return setting.Setting{
		ID:    db.ID,
		Key:   db.Key,
		Value: db.Value,
	}
}
