package dbmodel

import (
	rolepermission "final_project/internal/domain/role_permission"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"unique;size:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Users           []User           `gorm:"foreignKey:RoleID"`
	RolePermissions []RolePermission `gorm:"foreignKey:RoleID"`
}

// Domain -> DB
func RoleDomainToDB(domainPer rolepermission.Role) Role {
	return Role{
		ID:   domainPer.ID,
		Name: domainPer.Name,
	}
}

// DB -> Domain
func RoleDBToDomain(db Role) rolepermission.Role {
	return rolepermission.Role{
		ID:   db.ID,
		Name: db.Name,
		// Các trường CreatedAt, UpdatedAt, DeletedAt sẽ để GORM xử lý
	}
}
