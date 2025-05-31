package persistence

import (
	"context"
	"errors"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type RolePerRepoDB struct {
	db *gorm.DB
}

func NewRolePerRepoDB(db *gorm.DB) *RolePerRepoDB {
	return &RolePerRepoDB{db: db}
}

func (r *RolePerRepoDB) GetRoleNameByID(ctx context.Context, roleID uint) (string, error) {
	var roleName string

	if err := r.db.Debug().WithContext(ctx).Model(&rolepermission.Role{}).Select("name").Where("id = ?", roleID).Scan(&roleName).Error; err != nil {
		return "", err
	}

	return roleName, nil
}

func (r *RolePerRepoDB) GetRoleIDByName(ctx context.Context, roleName string) (uint, error) {
	var roleID uint

	if err := r.db.Debug().WithContext(ctx).Model(&rolepermission.Role{}).Select("id").Where("name = ?", roleName).Scan(&roleID).Error; err != nil {
		return 0, err
	}

	return roleID, nil
}

func (r *RolePerRepoDB) SavePermission(permissions *[]rolepermission.Permission) error {
	var dbPermission []dbmodel.Permission

	for _, value := range *permissions {
		dbPermission = append(dbPermission, dbmodel.PermissionDomainToDB(value))
	}

	if err := r.db.Debug().Model(&dbmodel.Permission{}).Create(&dbPermission).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) SaveRole(roles *[]rolepermission.Role) error {
	var dbRole []dbmodel.Role

	for _, value := range *roles {
		dbRole = append(dbRole, dbmodel.RoleDomainToDB(value))
	}

	if err := r.db.Debug().Model(&dbmodel.Role{}).Create(&dbRole).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) GetAllRoles(roles *[]rolepermission.Role) error {
	if err := r.db.Debug().Find(&roles).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) GetAllPermission(permissions *[]rolepermission.Permission) error {
	if err := r.db.Debug().Find(&permissions).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) SaveRolePermission(rolePermissions *[]rolepermission.RolePermission) error {
	var dbRolePer []dbmodel.RolePermission

	for _, value := range *rolePermissions {
		dbRolePer = append(dbRolePer, dbmodel.RolePerDomainToDB(value))
	}

	if err := r.db.Debug().Model(&dbmodel.RolePermission{}).Create(&dbRolePer).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePerRepoDB) IsRoleExisted(ctx context.Context, roleID uint) (bool, error) {
	var count int64 = 0

	if err := r.db.Debug().WithContext(ctx).Model(&rolepermission.Role{}).Where("id LIKE ?", roleID).Count(&count).Error; err != nil {
		return false, errors.New("Lỗi khi kiểm tra chức vụ đã tồn tại: " + err.Error())
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *RolePerRepoDB) IsRoleTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&rolepermission.Role{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *RolePerRepoDB) IsPermissionTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&rolepermission.Permission{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *RolePerRepoDB) IsRolePermissionTableEmpty(ctx context.Context) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&rolepermission.RolePermission{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
