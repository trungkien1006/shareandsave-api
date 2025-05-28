package request

import "context"

type Repository interface {
	Create(ctx context.Context, request *Request) error
}
