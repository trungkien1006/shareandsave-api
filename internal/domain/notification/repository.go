package notification

import "context"

type Repository interface {
	GetAll(ctx context.Context, notis *[]Notification, req GetAllNotiRequest, userID uint) (int, error)
	Create(ctx context.Context, noti *Notification) error
}
