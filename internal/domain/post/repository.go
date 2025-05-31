package post

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	AdminGetAll(ctx context.Context, posts *[]Post, filter filter.FilterRequest) (int, error)
	GetByID(ctx context.Context, post *Post, postID uint) error
	Save(ctx context.Context, post *CreatePost) error
	Update(ctx context.Context, post *Post) error
}
