package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type Notification struct {
    ID         uint           `gorm:"primaryKey;autoIncrement"`
    SenderID   uint           `gorm:"index"`
    ReceiverID uint           `gorm:"index"`
    Title      string         `gorm:"size:255"`
    Type       string         `gorm:"size:64"`
    TargetType string         `gorm:"size:32"`
    TargetID   uint
    Content    string         `gorm:"size:255"`
    IsRead     bool
    CreatedAt  time.Time
    UpdatedAt  time.Time
    DeletedAt  gorm.DeletedAt `gorm:"index"`
    // Relations
    Sender   User `gorm:"foreignKey:SenderID"`
    Receiver User `gorm:"foreignKey:ReceiverID"`
}