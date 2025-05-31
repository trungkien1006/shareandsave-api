package post

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	AdminGetAll(ctx context.Context, posts *[]AdminPost, filter filter.FilterRequest) (int, error)
	Save(ctx context.Context, post *Post) error
}
