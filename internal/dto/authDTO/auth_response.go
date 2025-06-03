package authdto

import userdto "final_project/internal/dto/userDTO"

type LoginResponseWrapper struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    LoginResponse `json:"data"`
}

type LoginResponse struct {
	JWT          string                `json:"jwt"`
	RefreshToken string                `json:"refreshToken"`
	User         userdto.CommonUserDTO `json:"user"`
}

type LogoutResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
