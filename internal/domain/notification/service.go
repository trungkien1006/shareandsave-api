package notification

import "context"

type Service interface {
	CreateAndPushSocket(ctx context.Context, noti *Notification) error
}
