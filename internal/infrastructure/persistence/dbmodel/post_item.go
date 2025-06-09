package dbmodel

type PostItem struct {
	ID       uint   `gorm:"primaryKey"`
	PostID   uint   `gorm:"type:int"`
	ItemID   uint   `gorm:"type:int"`
	Quantity int    `gorm:"type:INT"`
	Image    string `gorm:"type:LONGTEXT"`

	Post Post `gorm:"foreignKey:PostID"`
	Item Item `gorm:"foreignKey:ItemID"`
}
