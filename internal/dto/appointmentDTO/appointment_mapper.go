package appointmentdto

import "final_project/internal/domain/appointment"

// Domain to DTO
func AppointmentDomainToDTO(domain appointment.Appointment) AppointmentDTO {
	var items []AppointmentItemDTO

	for _, value := range domain.AppointmentItems {
		items = append(items, AppointmentItemDTO{
			ID:              value.ID,
			AppointmentID:   value.AppointmentID,
			ItemID:          value.ItemID,
			ActualQuantity:  value.ActualQuantity,
			MissingQuantity: value.MissingQuantity,
		})
	}

	return AppointmentDTO{
		ID:               domain.ID,
		UserID:           domain.UserID,
		StartTime:        domain.StartTime,
		EndTime:          domain.EndTime,
		Status:           domain.Status,
		AppointmentItems: items,
		CreatedAt:        domain.CreatedAt,
	}
}
