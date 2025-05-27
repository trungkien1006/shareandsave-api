package persistence

import (
	"context"
	"final_project/internal/domain/role_permission"

	"gorm.io/gorm"
)

type RolePerRepoDB struct {
	db *gorm.DB
}

func NewRolePerRepoDB(db *gorm.DB) *RolePerRepoDB {
	return &RolePerRepoDB{db: db}
}

func (r *RolePerRepoDB) SavePermission(permissions *[]role_permission.Permission) error {
	if err := r.db.Debug().Create(&permissions).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) SaveRole(roles *[]role_permission.Role) error {
	if err := r.db.Debug().Create(&roles).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) GetAllRoles(roles *[]role_permission.Role) error {
	if err := r.db.Debug().Find(&roles).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) GetAllPermission(permissions *[]role_permission.Permission) error {
	if err := r.db.Debug().Find(&permissions).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) SaveRolePermission(rolePermissions *[]role_permission.RolePermission) error {
	if err := r.db.Debug().Create(&rolePermissions).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) IsRoleTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&role_permission.Role{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *RolePerRepoDB) IsPermissionTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&role_permission.Permission{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *RolePerRepoDB) IsRolePermissionTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&role_permission.RolePermission{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
