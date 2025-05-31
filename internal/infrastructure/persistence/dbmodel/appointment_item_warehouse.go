package dbmodel

type AppointmentItemWarehouse struct {
	AppointmentID   uint `gorm:"primaryKey"`
	ItemWarehouseID uint `gorm:"primaryKey"`

	Appointment   Appointment   `gorm:"foreignKey:AppointmentID"`
	ItemWarehouse ItemWarehouse `gorm:"foreignKey:ItemWarehouseID"`
}
