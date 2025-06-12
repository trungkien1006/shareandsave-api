package dbmodel

import (
	"final_project/internal/domain/interest"
	"time"

	"gorm.io/gorm"
)

type Interest struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint `gorm:"index"`
	PostID    uint `gorm:"index"`
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:UserID"`
	Post Post `gorm:"foreignKey:PostID"`

	Comments     []Comment     `gorm:"foreignKey:InterestID"`
	Transactions []Transaction `gorm:"foreignKey:InterestID"`
}

// Domain to DB
func CreateDomainToDB(domain interest.Interest) Interest {
	return Interest{
		UserID: domain.UserID,
		PostID: domain.PostID,
		Status: domain.Status,
	}
}

// DB to Domain
func GetDTOToDomain(db Post) interest.PostInterest {
	domainItems := make([]interest.PostInterestItem, 0)
	domainInterest := make([]interest.Interest, 0)

	for _, value := range db.PostItem {
		domainItems = append(domainItems, interest.PostInterestItem{
			ID:              value.ID,
			ItemID:          value.ItemID,
			Name:            value.Item.Name,
			CategoryName:    value.Item.Category.Name,
			Quantity:        value.Quantity,
			CurrentQuantity: value.CurrentQuantity,
			Image:           value.Image,
		})
	}

	for _, value := range db.Interests {
		domainInterest = append(domainInterest, interest.Interest{
			ID:         value.ID,
			UserID:     value.UserID,
			UserName:   value.User.FullName,
			UserAvatar: value.User.Avatar,
			PostID:     value.PostID,
			Status:     value.Status,
			CreatedAt:  value.CreatedAt,
		})
	}

	return interest.PostInterest{
		ID:           db.ID,
		AuthorID:     db.AuthorID,
		AuthorName:   db.Author.FullName,
		AuthorAvatar: db.Author.Avatar,
		Title:        db.Title,
		Description:  db.Description,
		Slug:         db.Slug,
		Type:         db.Type,
		Items:        domainItems,
		Interests:    domainInterest,
		UpdatedAt:    db.UpdatedAt,
	}
}
