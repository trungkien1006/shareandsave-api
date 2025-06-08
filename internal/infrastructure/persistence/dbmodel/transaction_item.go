package dbmodel

type TransactionItem struct {
	TransactionID uint `gorm:"primaryKey"`
	ItemID        uint `gorm:"primaryKey"`
	Quantity      int  `gorm:"type:int"`

	Transaction Transaction `gorm:"foreignKey:TransactionID"`
	Item        Item        `gorm:"foreignKey:ItemID"`
}
