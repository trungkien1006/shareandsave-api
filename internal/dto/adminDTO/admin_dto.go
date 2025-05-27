package admindto

import "final_project/internal/domain/admin"

type AdminDTO struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullName"`
	Status   int8   `json:"status"`
	RoleName string `json:"roleName"`
}

func ToAdminDTO(a admin.Admin) AdminDTO {
	return AdminDTO{
		ID:       a.ID,
		Email:    a.Email,
		Fullname: a.FullName,
		Status:   a.Status,
		RoleName: a.Role.Name,
	}
}
