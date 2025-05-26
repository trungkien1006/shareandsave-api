package admindto

import "final_project/internal/domain/admin"

type AdminDTO struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Status   int8   `json:"status"`
	RoleID   uint   `json:"roleId"`
}

func ToAdminDTO(a admin.Admin) AdminDTO {
	return AdminDTO{
		ID:       a.ID,
		Email:    a.Email,
		Fullname: a.Fullname,
		Status:   a.Status,
		RoleID:   a.RoleID,
	}
}
