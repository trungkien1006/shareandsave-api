package notification

import "context"

type Repository interface {
	GetAllClient(ctx context.Context, notis *[]Notification, req GetAllNotiRequest, userID uint) (int64, int, error)
	GetAllAdmin(ctx context.Context, notis *[]Notification, req GetAllNotiRequest, userID uint) (int64, int, error)
	Create(ctx context.Context, noti *Notification) error
	ReadNoti(ctx context.Context, notiID uint) error
	ReadAllNoti(ctx context.Context, userID uint) error
}

type Notifier interface {
	SendToToken(token, title, body string, noti map[string]string) error
}
