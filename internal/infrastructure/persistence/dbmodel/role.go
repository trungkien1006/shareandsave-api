package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type Role struct {
    ID        uint           `gorm:"primaryKey;autoIncrement"`
    Name      string         `gorm:"unique;size:64"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    // Relations
    Admins []Admin `gorm:"foreignKey:RoleID"`
    RolePermissions []RolePermission `gorm:"foreignKey:RoleID"`
}