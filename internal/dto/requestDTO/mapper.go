package requestdto

import (
	"final_project/internal/domain/request"
	"final_project/internal/pkg/enums"
)

// DTO → Domain
func ToDomainRequest(dto CreateRequestSendOldItem) request.SendRequest {
	return request.SendRequest{
		UserID:              dto.UserID,
		Type:                enums.RequestType(dto.Type),
		Description:         dto.Description,
		AppointmentTime:     dto.AppointmentTime,
		AppointmentLocation: dto.AppointmentLocation,
		IsAnonymous:         dto.IsAnonymous,
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
		IsAnonymous:         domain.IsAnonymous,
	}
}
