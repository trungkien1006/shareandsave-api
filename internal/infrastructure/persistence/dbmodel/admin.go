package dbmodel

import (
	"final_project/internal/domain/admin"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"unique;size:255"`
	Password  string `gorm:"size:255"`
	FullName  string `gorm:"size:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Status    int8           `gorm:"type:TINYINT"`
	RoleID    uint           `gorm:"index"`
	// Relations
	Role           Role            `gorm:"foreignKey:RoleID"`
	ImportInvoices []ImportInvoice `gorm:"foreignKey:AdminID"`
	ExportInvoices []ExportInvoice `gorm:"foreignKey:AdminID"`
}

// Domain → DB
func AdminDomainToDB(a admin.Admin) Admin {
	return Admin{
		ID:       a.ID,
		Email:    a.Email,
		Password: a.Password,
		FullName: a.FullName,
		Status:   a.Status,
		RoleID:   a.RoleID,
	}
}

// DB → Domain
func AdminDBToDomain(a Admin) admin.Admin {
	return admin.Admin{
		ID:       a.ID,
		Email:    a.Email,
		Password: a.Password,
		FullName: a.FullName,
		Status:   a.Status,
		RoleID:   a.RoleID,
	}
}
