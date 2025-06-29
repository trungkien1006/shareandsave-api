package notificationdto

import "final_project/internal/domain/notification"

// Domain to DTO
func NotificationDomainToDTO(domain notification.Notification) NotificationDTO {
	return NotificationDTO{
		ID:         domain.ID,
		SenderID:   domain.SenderID,
		ReceiverID: domain.ReceiverID,
		Type:       domain.Type,
		TargetType: domain.TargetType,
		TargetID:   domain.TargetID,
		Content:    domain.Content,
		IsRead:     domain.IsRead,
		CreatedAt:  domain.CreatedAt,
	}
}
