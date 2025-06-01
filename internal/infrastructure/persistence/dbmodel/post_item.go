package dbmodel

type PostItem struct {
	PostID uint `gorm:"primaryKey"`
	ItemID uint `gorm:"primaryKey"`

	Post Post `gorm:"foreignKey:PostID"`
	Item Item `gorm:"foreignKey:ItemID"`
}
