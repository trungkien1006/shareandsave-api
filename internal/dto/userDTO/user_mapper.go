package userdto

import "final_project/internal/domain/user"

//Domain -> DTO
func DomainCommonUserToDTO(u user.User) CommonUserDTO {
	return CommonUserDTO{
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

//Domain -> DTO
func DomainAdminUserToDTO(u user.User) AdminUserDTO {
	var permissions []Permission

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
			Fullname:    u.FullName,
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
func DomainUserToDTO(u user.User) UserDTO {
	return UserDTO{
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
