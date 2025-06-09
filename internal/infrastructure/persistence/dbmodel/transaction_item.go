package dbmodel

type TransactionItem struct {
	TransactionID uint `gorm:"primaryKey"`
	PostItemID    uint `gorm:"primaryKey"`
	Quantity      int  `gorm:"type:int"`

	Transaction Transaction `gorm:"foreignKey:TransactionID"`
	PostItem    PostItem    `gorm:"foreignKey:PostItemID"`
}
