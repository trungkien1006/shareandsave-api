package admindto

import "final_project/internal/domain/admin"

// DTO → Domain
func ToDomainAdmin(dto CreateAdminRequest) admin.Admin {
	return admin.Admin{
		Email:    dto.Email,
		Password: dto.Password,
		FullName: dto.FullName,
		Status:   int8(dto.Status),
		RoleID:   dto.RoleID,
	}
}

// Domain → DTO (response)
func ToAdminDTO(domain admin.Admin, roleName string) AdminDTO {
	return AdminDTO{
		ID:       domain.ID,
		Email:    domain.Email,
		Fullname: domain.FullName,
		Status:   domain.Status,
		RoleName: roleName,
	}
}
