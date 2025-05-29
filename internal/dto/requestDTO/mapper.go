package requestdto

import (
	"final_project/internal/domain/request"
	"final_project/internal/pkg/enums"
)

// DTO → Domain
func DTOToDomainRequest(dto RequestSendOldItem) request.SendRequest {
	return request.SendRequest{
		ID:                  dto.ID,
		UserID:              dto.UserID,
		Type:                int(dto.Type),
		Description:         dto.Description,
		AppointmentTime:     dto.AppointmentTime,
		AppointmentLocation: dto.AppointmentLocation,
	}
}

// Domain → DTO (response)
func DomainToDTORequest(domain request.SendRequest) RequestSendOldItem {
	return RequestSendOldItem{
		ID:                  domain.ID,
		UserID:              domain.UserID,
		Type:                enums.RequestType(domain.Type),
		Description:         domain.Description,
		AppointmentTime:     domain.AppointmentTime,
		AppointmentLocation: domain.AppointmentLocation,
	}
}
