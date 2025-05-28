package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID          uint           `gorm:"primaryKey;autoIncrement"`
    Email       string         `gorm:"unique;size:255"`
    Password    string         `gorm:"size:255"`
    Avatar      string         `gorm:"type:LONGTEXT"`
    Active      bool
    FullName    string         `gorm:"size:64"`
    PhoneNumber string         `gorm:"unique;size:16"`
    Address     string         `gorm:"type:TEXT"`
    Status      int8           `gorm:"type:TINYINT"`
    GoodPoint   int            `gorm:"default:0"`
    Major       string         `gorm:"size:64"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
    // Relations
    PostsAuthored   []Post        `gorm:"foreignKey:AuthorID"`
    NotificationsSent     []Notification `gorm:"foreignKey:SenderID"`
    NotificationsReceived []Notification `gorm:"foreignKey:ReceiverID"`
}