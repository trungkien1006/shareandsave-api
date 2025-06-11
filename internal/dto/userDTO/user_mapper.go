package userdto

import "final_project/internal/domain/user"

//Domain -> DTO
func DomainCommonUserToDTO(u user.User) CommonUserDTO {
	return CommonUserDTO{
		ID:          u.ID,
		RoleID:      u.RoleID,
		RoleName:    u.RoleName,
		Email:       u.Email,
		FullName:    u.FullName,
		Avatar:      u.Avatar,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
		Status:      u.Status,
		GoodPoint:   u.GoodPoint,
		Major:       u.Major,
	}
}

//Domain -> DTO
func DomainAdminUserToDTO(u user.User) AdminUserDTO {
	permissions := make([]Permission, 0)

	for _, value := range u.Permissions {
		permissions = append(permissions, Permission{
			Code: value.Code,
		})
	}

	return AdminUserDTO{
		CommonUserDTO: CommonUserDTO{
			ID:          u.ID,
			RoleID:      u.RoleID,
			RoleName:    u.RoleName,
			Email:       u.Email,
			FullName:    u.FullName,
			Avatar:      u.Avatar,
			PhoneNumber: u.PhoneNumber,
			Address:     u.Address,
			Status:      u.Status,
			GoodPoint:   u.GoodPoint,
			Major:       u.Major,
		},
		Permissions: permissions,
	}
}

//Domain -> DTO
func DomainClientToDTO(u user.User) ClientDTO {
	return ClientDTO{
		ID:          u.ID,
		Email:       u.Email,
		Fullname:    u.FullName,
		Avatar:      u.Avatar,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
		Status:      u.Status,
		GoodPoint:   u.GoodPoint,
		Major:       u.Major,
	}
}

//Domain -> DTO
func DomainAdminToDTO(u user.User) AdminDTO {
	return AdminDTO{
		ID:          u.ID,
		RoleID:      u.RoleID,
		RoleName:    u.RoleName,
		Email:       u.Email,
		Fullname:    u.FullName,
		Avatar:      u.Avatar,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
		Status:      u.Status,
		GoodPoint:   u.GoodPoint,
		Major:       u.Major,
	}
}
