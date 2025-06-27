package userdto

import (
	"final_project/internal/domain/user"
	usergooddeed "final_project/internal/domain/user_good_deed"
)

// Domain -> DTO
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

// Domain -> DTO
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

// Domain -> DTO
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
		CreatedAt:   u.CreatedAt,
	}
}

// Domain -> DTO
func DomainUpdateUserToDTO(u user.User) UpdateUserDTO {
	return UpdateUserDTO{
		ID:          u.ID,
		Fullname:    u.FullName,
		Avatar:      u.Avatar,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
		Status:      u.Status,
		GoodPoint:   u.GoodPoint,
		Major:       u.Major,
	}
}

// Domain -> DTO
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
		CreatedAt:   u.CreatedAt,
	}
}

// Domain -> DTO
func DomainUserGoodDeedToDTO(ug usergooddeed.UserGoodDeedDetail) UserGoodDeedDetail {
	return UserGoodDeedDetail{
		ID:            ug.ID,
		UserID:        ug.UserID,
		UserName:      ug.UserName,
		GoodDeedType:  ug.GoodDeedType,
		GoodPoint:     ug.GoodPoint,
		TransactionID: ug.TransactionID,
		CreatedAt:     ug.CreatedAt,
		Items:         ug.Items,
	}
}

// Domain to DTO
func UserGoodDeedDomainToDTO(domain user.UserRank) UserRankDTO {
	userGoodDeeds := make([]UserGoodDeedDTO, 0)
	userGoodDeedMap := make(map[int]int, 0)

	for _, value := range domain.GoodDeeds {
		userGoodDeedMap[value.GoodDeedType] += 1
	}

	for key, value := range userGoodDeedMap {
		userGoodDeeds = append(userGoodDeeds, UserGoodDeedDTO{
			GoodDeedType:  key,
			GoodDeedCount: value,
		})
	}

	return UserRankDTO{
		UserID:     domain.UserID,
		UserName:   domain.UserName,
		UserAvatar: domain.UserAvatar,
		Major:      domain.Major,
		GoodPoint:  domain.GoodPoint,
		GoodDeeds:  userGoodDeeds,
	}
}
