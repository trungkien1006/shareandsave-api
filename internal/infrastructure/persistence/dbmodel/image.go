package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type Image struct {
    ID        uint           `gorm:"primaryKey;autoIncrement"`
    Target    string         `gorm:"size:32"`
    TargetID  uint
    Image     string         `gorm:"type:LONGTEXT"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}