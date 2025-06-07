package post

import (
	"context"
)

type Repository interface {
	AdminGetAll(ctx context.Context, posts *[]Post, filter AdminPostFilterRequest, userID uint) (int, error)
	GetAll(ctx context.Context, posts *[]PostWithCount, filter PostFilterRequest) (int, error)
	GetDetailByID(ctx context.Context, post *DetailPost, postID uint) error
	GetDetailBySlug(ctx context.Context, post *DetailPost, postSlug string) error
	GetByID(ctx context.Context, post *Post, postID uint) error
	Save(ctx context.Context, post *CreatePost) error
	Update(ctx context.Context, post *Post) error
	IsTableEmpty(ctx context.Context) (bool, error)
	IsExist(ctx context.Context, postID uint) (bool, error)
}
