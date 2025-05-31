package post

import (
	"context"
	"final_project/internal/domain/filter"
)

type Repository interface {
	AdminGetAll(ctx context.Context, posts *[]Post, filter filter.FilterRequest) (int, error)
	Save(ctx context.Context, post *Post) error
}
