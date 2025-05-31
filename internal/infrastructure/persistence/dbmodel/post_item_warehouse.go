package dbmodel

type PostItemWarehouse struct {
	PostID          uint `gorm:"primaryKey"`
	ItemWarehouseID uint `gorm:"primaryKey"`

	Post          Post          `gorm:"foreignKey:PostID"`
	ItemWarehouse ItemWarehouse `gorm:"foreignKey:ItemWarehouseID"`
}
