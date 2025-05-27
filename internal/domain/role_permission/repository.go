package role_permission

import "context"

type Repository interface {
	SavePermission(permissions *[]Permission) error
	SaveRole(roles *[]Role) error
	GetAllRoles(roles *[]Role) error
	GetAllPermission(ermissions *[]Permission) error
	SaveRolePermission(rolePermissions *[]RolePermission) error
	IsRoleTableEmpty(ctx context.Context) (bool, error)
	IsPermissionTableEmpty(ctx context.Context) (bool, error)
	IsRolePermissionTableEmpty(ctx context.Context) (bool, error)
}
