package itemdto

import (
	"final_project/internal/domain/item"
	"final_project/internal/infrastructure/persistence/dbmodel"
)

// Domain to DB
func DomainItemToDTO(domain item.Item) dbmodel.Item {
	return dbmodel.Item{
		ID:          domain.ID,
		Image:       domain.Image,
		Description: domain.Description,
		CategoryID:  domain.CategoryID,
	}
}
