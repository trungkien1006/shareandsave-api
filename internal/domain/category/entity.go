package category

import "time"

type Category struct {
	ID        uint
	Name      string
	CreatedAt time.Time // Assuming CreatedAt is a string for simplicity, adjust as needed
}
