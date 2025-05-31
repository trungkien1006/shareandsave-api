package userdto

type CommonUserDTO struct {
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
