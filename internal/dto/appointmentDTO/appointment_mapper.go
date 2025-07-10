package appointmentdto

import (
	"final_project/internal/domain/appointment"
	"time"
)

// Domain to DTO
func AppointmentDomainToDTO(domain appointment.Appointment) AppointmentDTO {
	var items []AppointmentItemDTO

	for _, value := range domain.AppointmentItems {
		items = append(items, AppointmentItemDTO{
			ID:              value.ID,
			AppointmentID:   value.AppointmentID,
			ItemID:          value.ItemID,
			ItemName:        value.ItemName,
			ItemImage:       value.ItemImage,
			CategoryName:    value.CategoryName,
			ActualQuantity:  value.ActualQuantity,
			MissingQuantity: value.MissingQuantity,
		})
	}

	return AppointmentDTO{
		ID:               domain.ID,
		UserID:           domain.UserID,
		UserName:         domain.UserName,
		StartTime:        domain.StartTime,
		EndTime:          domain.EndTime,
		Status:           domain.Status,
		AppointmentItems: items,
		CreatedAt:        domain.CreatedAt,
	}
}

// DTO to Domain
func UpdateAppointmentDTOToDomain(dto UpdateAppointmentRequest) appointment.Appointment {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	return appointment.Appointment{
		StartTime: dto.StartTime.In(loc),
		EndTime:   dto.EndTime.In(loc),
		Status:    int(dto.Status),
	}
}
