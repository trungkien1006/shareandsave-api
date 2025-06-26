package item

import "time"

type Item struct {
	ID           uint
	CategoryID   uint
	CategoryName string
	Name         string
	Description  string
	Image        string
	MaxClaim     int
	CreatedAt    time.Time
}
