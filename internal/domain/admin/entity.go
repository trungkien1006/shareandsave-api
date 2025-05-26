package admin

import (
	"final_project/internal/domain/role_permission"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	Fullname  string `gorm:"size:64"`
	Status    int8   `gorm:"type:TINYINT"`
	RoleID    uint   `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt       `gorm:"index"`
	Role      role_permission.Role `gorm:"foreignKey:RoleID"`
}

func NewAdmin(email, password, fullname string, status int8, roleID uint) *Admin {
	return &Admin{
		Email:    email,
		Password: password,
		Fullname: fullname,
		Status:   status,
		RoleID:   roleID,
	}
}
