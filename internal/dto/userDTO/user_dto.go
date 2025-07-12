package userdto

import (
	"time"
)

type CommonUserDTO struct {
	ID          uint   `json:"id"`
	RoleID      uint   `json:"roleID"`
	RoleName    string `json:"roleName"`
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
	Avatar      string `json:"avatar,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Address     string `json:"address,omitempty"`
	Status      int8   `json:"status"`
	GoodPoint   int    `json:"goodPoint"`
	Major       string `json:"major,omitempty"`
}

type AdminUserDTO struct {
	CommonUserDTO
	Permissions []Permission
}

type Permission struct {
	Code string `json:"code"`
}

type ClientDTO struct {
	ID          uint      `json:"id"`
	Email       string    `json:"email"`
	Fullname    string    `json:"fullName"`
	Avatar      string    `json:"avatar,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	Address     string    `json:"address,omitempty"`
	Status      int8      `json:"status"`
	GoodPoint   int       `json:"goodPoint"`
	Major       string    `json:"major,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}

type AdminDTO struct {
	ID          uint      `json:"id"`
	RoleID      uint      `json:"roleID"`
	RoleName    string    `json:"roleName"`
	Email       string    `json:"email"`
	Fullname    string    `json:"fullName"`
	Avatar      string    `json:"avatar,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	Address     string    `json:"address,omitempty"`
	Status      int8      `json:"status"`
	GoodPoint   int       `json:"goodPoint"`
	Major       string    `json:"major,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UpdateUserDTO struct {
	ID          uint   `json:"id"`
	Fullname    string `json:"fullName"`
	Avatar      string `json:"avatar,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Address     string `json:"address,omitempty"`
	Status      int8   `json:"status"`
	GoodPoint   int    `json:"goodPoint"`
	Major       string `json:"major,omitempty"`
}

type UserGoodDeedDetail struct {
	ID            uint                       `json:"id"`
	UserID        uint                       `json:"userID"`
	UserName      string                     `json:"userName"`
	GoodDeedType  int                        `json:"goodDeedType"`
	GoodPoint     int                        `json:"goodPoint"`
	TransactionID *uint                      `json:"transactionID"`
	CreatedAt     time.Time                  `json:"createdAt"`
	Items         []DetailTransactionItemDTO `json:"items"`
}

type DetailTransactionItemDTO struct {
	ItemID     uint   `json:"itemID"`
	ItemName   string `json:"itemName"`
	ItemImage  string `json:"itemImage"`
	PostItemID uint   `json:"postItemID"`
	Quantity   int    `json:"quantity"`
}

type UserRankDTO struct {
	UserID     uint              `json:"userID"`
	UserName   string            `json:"userName"`
	UserAvatar string            `json:"userAvatar"`
	Major      string            `json:"major"`
	GoodPoint  int               `json:"goodPoint"`
	GoodDeeds  []UserGoodDeedDTO `json:"goodDeeds"`
}

type UserGoodDeedDTO struct {
	GoodDeedType  int `json:"goodDeedType"`
	GoodDeedCount int `json:"goodDeedCount"`
}

type UserReportDTO struct {
	ID        uint              `json:"id"`
	FullName  string            `json:"fullName"`
	Major     string            `json:"major"`
	GoodPoint uint              `json:"goodPoint"`
	GoodDeeds []UserGoodDeedDTO `json:"goodDeeds"`
}
