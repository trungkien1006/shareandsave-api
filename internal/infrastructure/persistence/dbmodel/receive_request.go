package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type ReceiveRequest struct {
	ID                  uint           `gorm:"primaryKey;autoIncrement"` // id int [pk, increment]
	UserID              uint           `gorm:"not null"`                 // user_id int [ref: > user.id]
	Type                int            `gorm:"not null"`                 // type int
	Description         string         `gorm:"type:text"`                // description text
	Status              int8           `gorm:"type:tinyint;not null"`    // status tinyint
	ItemWarehouseID     *uint          `gorm:"default:null"`             // item_warehouse_id int [null, ref]
	PostID              *uint          `gorm:"default:null"`             // post_id int [null, ref]
	ReplyMessage        string         `gorm:"type:text"`                // reply_message text
	AppointmentTime     time.Time      `gorm:"type:datetime"`            // appointment_time datetime
	AppointmentLocation string         `gorm:"type:varchar(255)"`        // appointment_location varchar(255)
	CreatedAt           time.Time      `gorm:"autoCreateTime"`           // created_at timestamp
	UpdatedAt           time.Time      `gorm:"autoUpdateTime"`           // updated_at timestamp
	DeletedAt           gorm.DeletedAt `gorm:"index"`                    // deleted_at timestamp

	// Relations
	User          User           `gorm:"foreignKey:UserID"`          // Quan hệ với bảng user
	ItemWarehouse *ItemWarehouse `gorm:"foreignKey:ItemWarehouseID"` // Quan hệ với item_warehouse
	Post          *Post          `gorm:"foreignKey:PostID"`          // Quan hệ với post
}
