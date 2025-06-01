package itemdto

import "final_project/internal/domain/item"

type ItemDTO struct {
	ID          uint   `json:"id"`
	CategoryID  uint   `json:"categoryID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func ToItemDTO(item item.Item) ItemDTO {
	return ItemDTO{
		ID:          item.ID,
		CategoryID:  item.CategoryID,
		Name:        item.Name,
		Description: item.Description,
		Image:       item.Image,
	}
}
