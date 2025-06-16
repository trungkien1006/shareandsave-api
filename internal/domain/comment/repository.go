package comment

import "context"

type Repository interface {
	GetAll(ctx context.Context, domainComment *[]Comment, req GetComment) error
}
