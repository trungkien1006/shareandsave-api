package rolepermission

import "context"

type Repository interface {
	GetRoleNameByID(ctx context.Context, roleID uint) (string, error)
	GetRoleIDByName(ctx context.Context, roleName string) (uint, error)
	SavePermission(permissions *[]Permission) error
	SaveRole(roles *[]Role) error
	GetAllRoles(ctx context.Context, roles *[]Role) error
	GetAllPermission(ermissions *[]Permission) error
	SaveRolePermission(rolePermissions *[]RolePermission) error
	IsRoleExisted(ctx context.Context, roleID uint) (bool, error)
	IsRoleTableEmpty(ctx context.Context) (bool, error)
	IsPermissionTableEmpty(ctx context.Context) (bool, error)
	IsRolePermissionTableEmpty(ctx context.Context) (bool, error)
}
