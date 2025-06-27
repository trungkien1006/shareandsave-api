package user

import "time"

type User struct {
	ID          uint
	RoleID      uint
	RoleName    string
	Email       string
	Password    string
	Avatar      string
	Active      bool
	FullName    string
	PhoneNumber string
	Address     string
	Status      int8
	GoodPoint   int
	Major       string
	Permissions []Permission `gorm:"-"`
	CreatedAt   time.Time
}

type Permission struct {
	Code string
}

type UserRank struct {
	UserID     uint
	UserName   string
	UserAvatar string
	Major      string
	GoodPoint  int
	GoodDeeds  []UserGoodDeed `gorm:"-"`
}

type UserGoodDeed struct {
	GoodDeedType  int
	GoodDeedCount int
}
