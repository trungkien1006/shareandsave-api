package interestdto

import (
	"final_project/internal/domain/interest"
	"final_project/internal/pkg/enums"
)

// DTO to Domain
func CreateDTOToDomain(dto CreateInterest, userID uint) interest.Interest {
	return interest.Interest{
		PostID: dto.PostID,
		UserID: userID,
		Status: int(enums.InterestStatusStart),
	}
}

// DTO to Domain
func GetDTOToDomain(dto GetInterest) interest.GetInterest {
	return interest.GetInterest{
		Page:   dto.Page,
		Limit:  dto.Limit,
		Type:   int(dto.Type),
		Sort:   dto.Sort,
		Order:  dto.Order,
		Search: dto.Search,
	}
}

// Domain to DTO
func GetDomainToDTO(domain interest.PostInterest) PostInterest {
	domainItems := make([]PostInterestItem, 0)
	domainInterest := make([]Interest, 0)

	for _, value := range domain.Items {
		domainItems = append(domainItems, PostInterestItem{
			ID:              value.ID,
			ItemID:          value.ItemID,
			Name:            value.Name,
			CategoryName:    value.CategoryName,
			Quantity:        value.Quantity,
			CurrentQuantity: value.CurrentQuantity,
			Image:           value.Image,
		})
	}

	for _, value := range domain.Interests {
		domainInterest = append(domainInterest, Interest{
			ID:                 value.ID,
			UserID:             value.UserID,
			UserName:           value.UserName,
			UserAvatar:         value.UserAvatar,
			PostID:             value.PostID,
			Status:             value.Status,
			UnreadMessageCount: value.UnreadMessageCount,
			NewMessage:         value.NewMessage,
			NewMessageIsRead:   value.NewMessageIsRead,
			MessageFromID:      value.MessageFromID,
			CreatedAt:          value.CreatedAt,
		})
	}

	return PostInterest{
		ID:                 domain.ID,
		AuthorID:           domain.AuthorID,
		AuthorName:         domain.AuthorName,
		AuthorAvatar:       domain.AuthorAvatar,
		Title:              domain.Title,
		Description:        domain.Description,
		Slug:               domain.Slug,
		Type:               enums.PostType(domain.Type),
		Items:              domainItems,
		Interests:          domainInterest,
		UnreadMessageCount: domain.UnreadMessageCount,
		UpdatedAt:          domain.UpdatedAt,
		CreatedAt:          domain.CreatedAt,
	}
}
