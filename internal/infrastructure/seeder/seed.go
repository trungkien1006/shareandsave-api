package seeder

import (
	"context"
	"final_project/internal/domain/admin"
	"final_project/internal/domain/role_permission"
)

type Seeder struct {
	rolePerRepo role_permission.Repository
	adminRepo   admin.Repository
}

func (s *Seeder) NewSeeder(rolePerRepo role_permission.Repository) *Seeder {
	return &Seeder{
		rolePerRepo: rolePerRepo,
	}
}

func (s *Seeder) Seed() error {
	if err := s.seedPermission(); err != nil {
		return err
	}

	if err := s.seedRole(); err != nil {
		return err
	}

	if err := s.seedRolePer(); err != nil {
		return err
	}

	if err := s.seedAdmin(); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedPermission() error {
	var permissions = []role_permission.Permission{
		//user management permissions
		{Name: "Create User", Code: "create_user"},
		{Name: "Read User", Code: "read_user"},
		{Name: "Update User", Code: "update_user"},
		{Name: "Delete User", Code: "delete_user"},

		//item management permissions
		{Name: "Create Item", Code: "create_item"},
		{Name: "Read Item", Code: "read_item"},
		{Name: "Update Item", Code: "update_item"},
		{Name: "Delete Item", Code: "delete_item"},

		//post management permissions
		{Name: "Create Post", Code: "create_post"},
		{Name: "Read Post", Code: "read_post"},
		{Name: "Update Post", Code: "update_post"},
		{Name: "Delete Post", Code: "delete_post"},

		//admin management permissions
		{Name: "Create Admin", Code: "create_admin"},
		{Name: "Read Admin", Code: "read_admin"},
		{Name: "Update Admin", Code: "update_admin"},
		{Name: "Delete Admin", Code: "delete_admin"},

		//request management permissions
		{Name: "Read Request", Code: "read_request"},
		{Name: "Reply Request", Code: "reply_request"},
		{Name: "Delete Request", Code: "delete_request"},

		//notification management permissions
		{Name: "Create Notification", Code: "create_notification"},

		//import_invoice management permissions
		{Name: "Read Import Invoice", Code: "read_import_invoice"},
		{Name: "Create Import Invoice", Code: "read_import_invoice"},
		{Name: "Update Import Invoice", Code: "update_import_invoice"},
		{Name: "Lock Import Invoice", Code: "lock_import_invoice"},
		{Name: "Delete Import Invoice", Code: "delete_import_invoice"},

		//export_invoice management permissions
		{Name: "Read Export Invoice", Code: "read_export_invoice"},
		{Name: "Create Export Invoice", Code: "read_export_invoice"},
		{Name: "Update Export Invoice", Code: "update_export_invoice"},
		{Name: "Lock Export Invoice", Code: "lock_export_invoice"},
		{Name: "Delete Export Invoice", Code: "delete_export_invoice"},

		//warehouse management permissions
		{Name: "Read Warehouse", Code: "read_warehouse"},

		//item_warehouse management permissions
		{Name: "Read Item Warehouse", Code: "read_item_warehouse"},
	}

	if err := s.rolePerRepo.SavePermission(&permissions); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedRole() error {
	var roles = []role_permission.Role{
		{Name: "Super Admin"},
		{Name: "Content Manager"},
		{Name: "Warehouse Manager"},
		{Name: "Human Resources Manager"},
		{Name: "Client Manager"},
	}

	if err := s.rolePerRepo.SaveRole(&roles); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedRolePer() error {
	var rolePermissionConfig = map[string][]string{
		"Super Admin":             {"*"},
		"Content Manager":         {"create_post", "read_post", "update_post", "delete_post", "read_notification"},
		"Warehouse Manager":       {"read_warehouse", "read_item_warehouse", "read_import_invoice", "create_import_invoice", "update_import_invoice", "lock_import_invoice", "delete_import_invoice", "read_export_invoice", "create_export_invoice", "update_export_invoice", "lock_export_invoice", "delete_export_invoice", "read_notification"},
		"Human Resources Manager": {"read_admin", "update_admin", "delete_admin", "create_admin", "read_notification"},
		"Client Manager":          {"read_user", "update_user", "delete_user", "read_notification", "read_request", "reply_request", "delete_request"},
	}

	var roles []role_permission.Role
	var permissions []role_permission.Permission
	var rolePerms []role_permission.RolePermission

	if err := s.rolePerRepo.GetAllRoles(&roles); err != nil {
		return err
	}

	if err := s.rolePerRepo.GetAllPermission(&permissions); err != nil {
		return err
	}

	permCodeToID := make(map[string]uint)

	for _, p := range permissions {
		permCodeToID[p.Code] = p.ID
	}

	for _, role := range roles {
		codes, exists := rolePermissionConfig[role.Name]

		if !exists {
			continue // nếu role đó không có cấu hình thì bỏ qua
		}

		if len(codes) == 1 && codes[0] == "*" {
			for _, p := range permissions {
				rolePerms = append(rolePerms, role_permission.RolePermission{
					RoleID:       role.ID,
					PermissionID: p.ID,
				})
			}
		} else {
			for _, code := range codes {
				if permID, ok := permCodeToID[code]; ok {
					rolePerms = append(rolePerms, role_permission.RolePermission{
						RoleID:       role.ID,
						PermissionID: permID,
					})
				}
			}
		}
	}

	if err := s.rolePerRepo.SaveRolePermission(&rolePerms); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedAdmin() error {
	ctx := context.Background()

	var roles []role_permission.Role

	if err := s.rolePerRepo.GetAllRoles(&roles); err != nil {
		return err
	}

	roleMap := make(map[string]uint)
	for _, r := range roles {
		roleMap[r.Name] = r.ID
	}

	admins := []admin.Admin{
		*admin.NewAdmin("superadmin@example.com", "hashed_password_1", "Super Admin", 1, roleMap["Super Admin"]),
		*admin.NewAdmin("content@example.com", "hashed_password_2", "Content Manager", 1, roleMap["Content Manager"]),
		*admin.NewAdmin("warehouse@example.com", "hashed_password_3", "Warehouse Manager", 1, roleMap["Warehouse Manager"]),
		*admin.NewAdmin("hr@example.com", "hashed_password_4", "HR Manager", 1, roleMap["Human Resources Manager"]),
		*admin.NewAdmin("client@example.com", "hashed_password_5", "Client Manager", 1, roleMap["Client Manager"]),
	}

	for i := range admins {
		if err := s.adminRepo.Save(ctx, &admins[i]); err != nil {
			return err
		}
	}

	return nil
}
