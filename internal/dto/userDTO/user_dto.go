package userDTO

import "final_project/internal/domain/user"

type UserDTO struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Fullname    string `json:"fullName"`
	Avatar      string `json:"avatar,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Address     string `json:"address,omitempty"`
	Status      int8   `json:"status"`
	GoodPoint   int    `json:"goodPoint"`
	Major       string `json:"major,omitempty"`
}

func ToUserDTO(u user.User) UserDTO {
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
