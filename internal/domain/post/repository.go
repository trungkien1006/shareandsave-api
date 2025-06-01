package post

import (
	"context"
)

type Repository interface {
	AdminGetAll(ctx context.Context, posts *[]Post, filter PostFilterRequest) (int, error)
	GetByID(ctx context.Context, post *Post, postID uint) error
	Save(ctx context.Context, post *CreatePost) error
	Update(ctx context.Context, post *Post) error
}
