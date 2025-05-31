package dbmodel

import (
	rolepermission "final_project/internal/domain/role_permission"
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"unique;size:255"`
	Code      string `gorm:"unique;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	RolePermissions []RolePermission `gorm:"foreignKey:PermissionID"`
}

// Domain -> DB
func PermissionDomainToDB(domainPer rolepermission.Permission) Permission {
	return Permission{
		ID:   domainPer.ID,
		Name: domainPer.Name,
		Code: domainPer.Code,
	}
}

// DB -> Domain
func PermissionDBToDomain(db Permission) rolepermission.Permission {
	return rolepermission.Permission{
		ID:   db.ID,
		Name: db.Name,
		Code: db.Code,
		// Các trường CreatedAt, UpdatedAt, DeletedAt sẽ để GORM xử lý
	}
}
