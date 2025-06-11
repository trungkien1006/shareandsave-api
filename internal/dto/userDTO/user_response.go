package userdto

// Get client
type GetClientResponseWrapper struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    GetClientResponse `json:"data"`
}

type GetClientResponse struct {
	Clients   []ClientDTO `json:"clients"`
	TotalPage int         `json:"totalPage"`
}

// Get admin
type GetAdminResponseWrapper struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    GetAdminResponse `json:"data"`
}

type GetAdminResponse struct {
	Admins    []AdminDTO `json:"admins"`
	TotalPage int        `json:"totalPage"`
}

// Get client by id
type GetClientByIDResponseWrapper struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    GetClientByIDResponse `json:"data"`
}

type GetClientByIDResponse struct {
	Client ClientDTO `json:"client"`
}

// Get admin by id
type GetAdminByIDResponseWrapper struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    GetAdminByIDResponse `json:"data"`
}

type GetAdminByIDResponse struct {
	Admin AdminDTO `json:"admin"`
}

// Create client
type CreateClientResponseWrapper struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    CreateClientResponse `json:"data"`
}

type CreateClientResponse struct {
	Client ClientDTO `json:"client"`
}

// Create admin
type CreateAdminResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    CreateAdminResponse `json:"data"`
}

type CreateAdminResponse struct {
	Admin AdminDTO `json:"admin"`
}

// Update client
type UpdateClientResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Update admin
type UpdateAdminResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//Delete client
type DeleteClientResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//Delete admin
type DeleteAdminResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
