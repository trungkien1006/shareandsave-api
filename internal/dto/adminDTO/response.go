package admindto

type GetAdminResponseWrapper struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    GetAdminResponse `json:"data"`
}

type GetAdminResponse struct {
	Admins    []AdminDTO `json:"admins"`
	TotalPage int        `json:"totalPage"`
}

type GetAdminByIDResponseWrapper struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    GetAdminByIDResponse `json:"data"`
}

type GetAdminByIDResponse struct {
	Admin AdminDTO `json:"admin"`
}

type CreateAdminResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    CreateAdminResponse `json:"data"`
}

type CreateAdminResponse struct {
	Admin AdminDTO `json:"admin"`
}

type UpdateAdminResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DeleteAdminResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
