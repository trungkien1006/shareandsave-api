package dbmodel

import (
	"final_project/internal/domain/user"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	RoleID      uint   `gorm:"index"`
	Email       string `gorm:"unique;size:255"`
	PhoneNumber string `gorm:"unique;size:16"`
	Password    string `gorm:"size:255"`
	Avatar      string `gorm:"type:LONGTEXT"`
	Active      bool
	FullName    string `gorm:"size:64"`
	Address     string `gorm:"type:TEXT"`
	Status      int8   `gorm:"type:TINYINT"`
	GoodPoint   int    `gorm:"default:0"`
	Major       string `gorm:"size:64"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// n-1: User thuộc về Role
	Role Role `gorm:"foreignKey:RoleID"`

	// 1-n: User có nhiều post, interest, comment, transaction, appointment, notification
	Posts             []Post         `gorm:"foreignKey:AuthorID"`
	Interests         []Interest     `gorm:"foreignKey:UserID"`
	CommentsSent      []Comment      `gorm:"foreignKey:SenderID"`
	CommentsRecv      []Comment      `gorm:"foreignKey:ReceiverID"`
	TransactionsSent  []Transaction  `gorm:"foreignKey:SenderID"`
	TransactionsRecv  []Transaction  `gorm:"foreignKey:ReceiverID"`
	Appointments      []Appointment  `gorm:"foreignKey:UserID"`
	NotificationsSent []Notification `gorm:"foreignKey:SenderID"`
	NotificationsRecv []Notification `gorm:"foreignKey:ReceiverID"`
}

func ToDomainUser(dbUser User) user.User {
	permissions := make([]user.Permission, 0)

	if dbUser.Role.RolePermissions != nil {
		for _, value := range dbUser.Role.RolePermissions {
			permissions = append(permissions, user.Permission{
				Code: value.Permission.Code,
			})
		}
	}

	return user.User{
		ID:          dbUser.ID,
		RoleID:      dbUser.RoleID,
		RoleName:    dbUser.Role.Name, // map từ quan hệ Role
		Email:       dbUser.Email,
		Password:    dbUser.Password,
		Avatar:      dbUser.Avatar,
		Active:      dbUser.Active,
		FullName:    dbUser.FullName,
		PhoneNumber: dbUser.PhoneNumber,
		Address:     dbUser.Address,
		Status:      dbUser.Status,
		GoodPoint:   dbUser.GoodPoint,
		Major:       dbUser.Major,
		Permissions: permissions,
	}
}

func ToDBUser(u user.User) User {
	return User{
		ID:          u.ID,
		RoleID:      u.RoleID,
		Email:       u.Email,
		Password:    u.Password,
		Avatar:      u.Avatar,
		Active:      u.Active,
		FullName:    u.FullName,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
		Status:      u.Status,
		GoodPoint:   u.GoodPoint,
		Major:       u.Major,
	}
}
