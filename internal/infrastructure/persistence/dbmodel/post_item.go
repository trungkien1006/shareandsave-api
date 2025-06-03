package dbmodel

type PostItem struct {
	PostID   uint   `gorm:"primaryKey"`
	ItemID   uint   `gorm:"primaryKey"`
	Quantity int    `gorm:"type:INT"`
	Image    string `gorm:"type:LONGTEXT"`

	Post Post `gorm:"foreignKey:PostID"`
	Item Item `gorm:"foreignKey:ItemID"`
}
