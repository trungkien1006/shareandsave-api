package notification

import "context"

type Repository interface {
	GetAllClient(ctx context.Context, notis *[]Notification, req GetAllNotiRequest, userID uint) (int, error)
	GetAllAdmin(ctx context.Context, notis *[]Notification, req GetAllNotiRequest, userID uint) (int, error)
	Create(ctx context.Context, noti *Notification) error
}
