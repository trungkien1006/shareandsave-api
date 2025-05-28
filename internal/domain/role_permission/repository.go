package rolepermission

import "context"

type Repository interface {
	GetRoleNameByID(ctx context.Context, roleID uint) (string, error)
	SavePermission(permissions *[]Permission) error
	SaveRole(roles *[]Role) error
	GetAllRoles(roles *[]Role) error
	GetAllPermission(ermissions *[]Permission) error
	SaveRolePermission(rolePermissions *[]RolePermission) error
	IsRoleTableEmpty(ctx context.Context) (bool, error)
	IsPermissionTableEmpty(ctx context.Context) (bool, error)
	IsRolePermissionTableEmpty(ctx context.Context) (bool, error)
}
