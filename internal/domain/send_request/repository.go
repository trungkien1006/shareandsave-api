package sendrequest

import "context"

type Repository interface {
	Create(ctx context.Context, request *SendRequest) error
}
