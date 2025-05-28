package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type Request struct {
    ID                  uint           `gorm:"primaryKey;autoIncrement"`
    UserID              uint           `gorm:"index"`
    RequestType         int
    Description         string         `gorm:"type:TEXT"`
    IsAnonymous         bool
    Status              int8           `gorm:"type:TINYINT"`
    ItemWarehouseID     *uint          `gorm:"index"`
    PostID              *uint          `gorm:"index"`
    ReplyMessage        string         `gorm:"type:TEXT"`
    AppointmentTime     time.Time
    AppointmentLocation string         `gorm:"size:255"`
    CreatedAt           time.Time
    UpdatedAt           time.Time
    DeletedAt           gorm.DeletedAt `gorm:"index"`
    // Relations
    User           User           `gorm:"foreignKey:UserID"`
    ItemWarehouse  *ItemWarehouse `gorm:"foreignKey:ItemWarehouseID"`
    Post           *Post          `gorm:"foreignKey:PostID"`
}