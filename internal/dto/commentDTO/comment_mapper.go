package commentdto

import "final_project/internal/domain/comment"

// Domain to DTO
func CommentDomainToDTO(domain comment.Comment) CommentDTO {
	return CommentDTO{
		ID:         domain.ID,
		InterestID: domain.InterestID,
		SenderID:   domain.SenderID,
		ReceiverID: domain.ReceiverID,
		Content:    domain.Content,
		IsRead:     domain.IsRead,
		CreatedAt:  domain.CreatedAt,
	}
}
