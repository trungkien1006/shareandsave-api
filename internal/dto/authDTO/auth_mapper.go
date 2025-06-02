package authdto

import "final_project/internal/domain/auth"

// DTO to Domain
func AuthDTOToDomain(dto LoginRequest) auth.AuthLogin {
	return auth.AuthLogin{
		Email:    dto.Email,
		Password: dto.Password,
		Device:   dto.Device,
	}
}
