package request

import (
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
	Status              int8   `gorm:"type:TINYINT"`
	ItemWarehouseID     uint   `gorm:"index"`
	PostID              uint   `gorm:"index"`
	ReplyMessage        string `gorm:"type:TEXT"`
	AppointmentTime     time.Time
	AppointmentLocation string `gorm:"size:255"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`

	User user.User `gorm:"foreignKey:UserID"`
	Post post.Post `gorm:"foreignKey:PostID"`
}

func NewRequest(userID uint, requestType int, description string, isAnonymous bool, item, itemWarehouseID, postID uint, appointmentTime time.Time, location string, status int8) *Request {
	return &Request{
		UserID:              userID,
		RequestType:         requestType,
		Description:         description,
		IsAnonymous:         isAnonymous,
		ItemWarehouseID:     itemWarehouseID,
		PostID:              postID,
		AppointmentTime:     appointmentTime,
		AppointmentLocation: location,
		Status:              status,
	}
}
