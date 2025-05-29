package dbmodel

import (
	"final_project/internal/domain/request"
	"time"

	"gorm.io/gorm"
)

type SendRequest struct {
	ID                  uint           `gorm:"primaryKey;autoIncrement"` // id int [pk, increment]
	UserID              uint           `gorm:"not null"`                 // user_id int [ref: > user.id]
	Type                int            `gorm:"not null"`                 // type int
	Description         string         `gorm:"type:text"`                // description text
	Status              int8           `gorm:"type:tinyint;not null"`    // status tinyint
	ReplyMessage        string         `gorm:"type:text"`                // reply_message text
	AppointmentTime     time.Time      `gorm:"type:datetime"`            // appointment_time datetime
	AppointmentLocation string         `gorm:"type:varchar(255)"`        // appointment_location varchar(255)
	CreatedAt           time.Time      `gorm:"autoCreateTime"`           // created_at timestamp
	UpdatedAt           time.Time      `gorm:"autoUpdateTime"`           // updated_at timestamp
	DeletedAt           gorm.DeletedAt `gorm:"index"`                    // deleted_at timestamp
	// Relations
	User User `gorm:"foreignKey:UserID"` // Quan hệ với bảng user
}

// Domain → DB
func RequestDomainToDB(a request.SendRequest) request.SendRequest {
	return request.SendRequest{
		ID:                  a.ID,
		UserID:              a.UserID,
		Type:                a.Type,
		Description:         a.Description,
		Status:              a.Status,
		ReplyMessage:        a.ReplyMessage,
		AppointmentTime:     a.AppointmentTime,
		AppointmentLocation: a.AppointmentLocation,
	}
}

// DB → Domain
func SendRequestToDomain(a request.SendRequest) request.SendRequest {
	return request.SendRequest{
		ID:                  a.ID,
		UserID:              a.UserID,
		Type:                a.Type,
		Description:         a.Description,
		Status:              a.Status,
		ReplyMessage:        a.ReplyMessage,
		AppointmentTime:     a.AppointmentTime,
		AppointmentLocation: a.AppointmentLocation,
	}
}
