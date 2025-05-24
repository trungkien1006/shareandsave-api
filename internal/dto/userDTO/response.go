package userDTO

import "final_project/internal/domain/user"

type UserDTO struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Fullname    string `json:"fullName"`
	Avatar      string `json:"avatar,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Address     string `json:"address,omitempty"`
	Status      int    `json:"status"`
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
	User UserDTO `json:"user"`
}

type CreateUserResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    CreateUserResponse `json:"data"`
}

type CreateUserResponse struct {
	User UserDTO `json:"user"`
}

type UpdateUserResponseWrapper struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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
	}
}
