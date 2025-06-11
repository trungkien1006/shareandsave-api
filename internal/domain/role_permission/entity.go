package rolepermission

type Role struct {
	ID   uint
	Name string
}

type Permission struct {
	ID   uint
	Name string
	Code string
}

type RolePermission struct {
	RoleID       uint
	PermissionID uint
}

type RolePermissionList struct {
	ID          uint
	Name        string
	Permissions []Permission
}
