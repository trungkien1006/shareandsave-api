package dbmodel

import (
    "time"
    "gorm.io/gorm"
)

type Post struct {
    ID          uint           `gorm:"primaryKey;autoIncrement"`
    AuthorID    uint           `gorm:"index"`
    ItemID      uint           `gorm:"index"`
    Title       string         `gorm:"size:255"`
    Description string         `gorm:"type:TEXT"`
    Status      int8           `gorm:"type:TINYINT"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
    // Relations
    Author User `gorm:"foreignKey:AuthorID"`
    Item   Item `gorm:"foreignKey:ItemID"`
    Requests []Request `gorm:"foreignKey:PostID"`
}