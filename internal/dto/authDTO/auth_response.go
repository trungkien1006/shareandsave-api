package authdto

import userdto "final_project/internal/dto/userDTO"

type ClientLoginResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    ClientLoginResponse `json:"data"`
}

type ClientLoginResponse struct {
	JWT          string                `json:"jwt"`
	RefreshToken string                `json:"refreshToken"`
	User         userdto.CommonUserDTO `json:"user"`
}

type AdminLoginResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    AdminLoginResponse `json:"data"`
}

type AdminLoginResponse struct {
	JWT          string               `json:"jwt"`
	RefreshToken string               `json:"refreshToken"`
	User         userdto.AdminUserDTO `json:"user"`
}

type LogoutResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetAccessTokenResponseWrapper struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    GetAccessTokenResponse `json:"data"`
}

type GetAccessTokenResponse struct {
	JWT string `json:"jwt"`
}

type AdminGetMeResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    AdminGetMeResponse `json:"data"`
}

type AdminGetMeResponse struct {
	User userdto.AdminUserDTO `json:"user"`
}

type ClientGetMeResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    ClientGetMeResponse `json:"data"`
}

type ClientGetMeResponse struct {
	User userdto.CommonUserDTO `json:"user"`
}
