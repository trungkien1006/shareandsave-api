package dbmodel

type TransactionItem struct {
	TransactionID uint `gorm:"primaryKey"`
	ItemID        uint `gorm:"primaryKey"`

	Transaction Transaction `gorm:"foreignKey:TransactionID"`
	Item        Item        `gorm:"foreignKey:ItemID"`
}
