package comment

import "context"

type Repository interface {
	GetAll(ctx context.Context, domainComment *[]Comment, req GetComment) error
	CreateBatch(ctx context.Context, domainComments *[]Comment) error
	Create(ctx context.Context, domainComment *Comment) error
	UpdateReadMessage(ctx context.Context, interestID uint) error
}
