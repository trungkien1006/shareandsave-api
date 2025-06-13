package userdto

type CommonUserDTO struct {
	ID          uint   `json:"id"`
	RoleID      uint   `json:"roleID"`
	RoleName    string `json:"roleName"`
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
	Avatar      string `json:"avatar,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Address     string `json:"address,omitempty"`
	Status      int8   `json:"status"`
	GoodPoint   int    `json:"goodPoint"`
	Major       string `json:"major,omitempty"`
}

type AdminUserDTO struct {
	CommonUserDTO
	Permissions []Permission
}

type Permission struct {
	Code string `json:"code"`
}

type ClientDTO struct {
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

type AdminDTO struct {
	ID          uint   `json:"id"`
	RoleID      uint   `json:"roleID"`
	RoleName    string `json:"roleName"`
	Email       string `json:"email"`
	Fullname    string `json:"fullName"`
	Avatar      string `json:"avatar,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Address     string `json:"address,omitempty"`
	Status      int8   `json:"status"`
	GoodPoint   int    `json:"goodPoint"`
	Major       string `json:"major,omitempty"`
}

type UpdateUserDTO struct {
	ID          uint   `json:"id"`
	Fullname    string `json:"fullName"`
	Avatar      string `json:"avatar,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Address     string `json:"address,omitempty"`
	Status      int8   `json:"status"`
	GoodPoint   int    `json:"goodPoint"`
	Major       string `json:"major,omitempty"`
}
