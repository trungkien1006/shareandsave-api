package request

import (
	"final_project/internal/domain/item"
	"final_project/internal/domain/post"
	"os/user"
	"time"

	"gorm.io/gorm"
)

type Request struct {
	ID                  uint `gorm:"primaryKey;autoIncrement"`
	UserID              uint `gorm:"index"`
	RequestType         int
	Description         string `gorm:"type:TEXT"`
	IsAnonymous         bool
	Status              int8 `gorm:"type:TINYINT"`
	Item                int
	ItemID              uint `gorm:"index"`
	PostID              uint `gorm:"index"`
	AppointmentTime     time.Time
	AppointmentLocation string `gorm:"size:255"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`

	User    user.User `gorm:"foreignKey:UserID"`
	ItemRef item.Item `gorm:"foreignKey:ItemID"`
	Post    post.Post `gorm:"foreignKey:PostID"`
}

func NewRequest(userID uint, requestType int, description string, isAnonymous bool, item, itemID, postID uint, appointmentTime time.Time, location string, status int8) *Request {
	return &Request{
		UserID:              userID,
		RequestType:         requestType,
		Description:         description,
		IsAnonymous:         isAnonymous,
		Item:                int(item),
		ItemID:              itemID,
		PostID:              postID,
		AppointmentTime:     appointmentTime,
		AppointmentLocation: location,
		Status:              status,
	}
}
