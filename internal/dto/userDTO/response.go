package userDTO

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
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DeleteUserResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
