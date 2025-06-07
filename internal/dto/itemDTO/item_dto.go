package itemdto

import (
	"final_project/internal/domain/item"
	"time"
)

type ItemDTO struct {
	ID           uint      `json:"id"`
	CategoryID   uint      `json:"categoryID"`
	CategoryName string    `json:"categoryName"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Image        string    `json:"image"`
	CreatedAt    time.Time `json:"createdAt"`
}

func ToItemDTO(item item.Item) ItemDTO {
	return ItemDTO{
		ID:           item.ID,
		CategoryID:   item.CategoryID,
		CategoryName: item.CategoryName,
		Name:         item.Name,
		Description:  item.Description,
		Image:        item.Image,
		CreatedAt:    item.CreatedAt,
	}
}
