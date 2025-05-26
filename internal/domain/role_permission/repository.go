package role_permission

type Repository interface {
	SavePermission(permissions *[]Permission) error
	SaveRole(roles *[]Role) error
	GetAllRoles(roles *[]Role) error
	GetAllPermission(ermissions *[]Permission) error
	SaveRolePermission(rolePermissions *[]RolePermission) error
}
