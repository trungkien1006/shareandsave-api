package admin

import (
	"final_project/internal/domain/role_permission"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint                 `gorm:"primaryKey;autoIncrement"`
	Email     string               `gorm:"unique;size:255;not null"`
	Password  string               `gorm:"size:255;not null"`
	FullName  string               `gorm:"size:64"`
	Status    int8                 `gorm:"type:TINYINT"`
	RoleID    uint                 `gorm:"index"`
	Role      role_permission.Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewAdmin(email, password, fullName string, status int8, roleID uint) *Admin {
	return &Admin{
		Email:    email,
		Password: password,
		FullName: fullName,
		Status:   status,
		RoleID:   roleID,
	}
}
