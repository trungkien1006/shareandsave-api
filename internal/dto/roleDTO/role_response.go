package roledto

// Response trả về danh sách role
type GetRoleResponse struct {
	Roles []RoleDTO `json:"roles"`
}

// Wrapper cho response danh sách Role
type GetRoleResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    GetRoleResponse `json:"data"`
}
