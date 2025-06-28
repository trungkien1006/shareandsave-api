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

// DTO to Domain
func SignUpAuthDTOToDomain(dto SignUpRequest) auth.AuthSignUp {
	return auth.AuthSignUp{
		VerifyToken: dto.VerifyToken,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		FullName:    dto.FullName,
		Password:    dto.Password,
		RePassword:  dto.RePassword,
	}
}

// DTO to Domain
func ResetPasswordDTOToDomain(dto ResetPasswordRequest) auth.AuthResetPassword {
	return auth.AuthResetPassword{
		VerifyToken: dto.VerifyToken,
		Email:       dto.Email,
		Password:    dto.Password,
		RePassword:  dto.RePassword,
	}
}
