package dbmodel

type RolePermission struct {
    RoleID       uint `gorm:"primaryKey"`
    PermissionID uint `gorm:"primaryKey"`
    // Relations
    Role       Role       `gorm:"foreignKey:RoleID"`
    Permission Permission `gorm:"foreignKey:PermissionID"`
}