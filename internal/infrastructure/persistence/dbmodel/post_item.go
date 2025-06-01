package dbmodel

type PostItem struct {
	PostID   uint `gorm:"primaryKey"`
	ItemID   uint `gorm:"primaryKey"`
	Quantity int

	Post Post `gorm:"foreignKey:PostID"`
	Item Item `gorm:"foreignKey:ItemID"`
}
