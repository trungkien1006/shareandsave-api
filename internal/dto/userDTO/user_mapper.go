package userdto

import "final_project/internal/domain/user"

//Domain -> DTO
func DomainToDTO(u user.User) UserDTO {
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
