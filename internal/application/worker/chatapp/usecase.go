package chatapp

import (
	"context"
	"final_project/internal/domain/comment"
)

type UseCase struct {
	commentRepo comment.Repository
}

func NewUseCase(commentRepo comment.Repository) *UseCase {
	return &UseCase{commentRepo: commentRepo}
}

func (uc *UseCase) CreateMessage(ctx context.Context, messages []map[string]string) error {

	return nil
}
