package admindto

type AdminDTO struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullName"`
	Status   int8   `json:"status"`
	RoleName string `json:"roleName"`
}
