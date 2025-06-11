package roledto

import (
	rolepermission "final_project/internal/domain/role_permission"
)

// Domain -> DTO
func RoleDomainToDTO(domain rolepermission.Role) RoleDTO {
	return RoleDTO{
		ID:   domain.ID,
		Name: domain.Name,
	}
}
