package categorydto

import "time"

type CategoryDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"` // Assuming CreatedAt is a string for simplicity, adjust as needed
}
