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
}

type GetUserResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    GetUserResponse `json:"data"`
}

type GetUserResponse struct {
	Users     []UserDTO `json:"users"`
	TotalPage int       `json:"totalPage"`
}

type GetUserByIDResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    GetUserByIDResponse `json:"data"`
}

type GetUserByIDResponse struct {
	User []UserDTO `json:"user"`
}

func ToUserDTO(u user.User) UserDTO {
	return UserDTO{
		ID:          u.ID,
		Email:       u.Email,
		Fullname:    u.Fullname,
		Avatar:      u.Avatar,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
		Status:      u.Status,
		GoodPoint:   u.GoodPoint,
	}
}
