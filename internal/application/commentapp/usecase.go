package commentapp

import "final_project/internal/domain/comment"

type UseCase struct {
	repo comment.Repository
}

func NewUseCase(r comment.Repository) *UseCase {
	return &UseCase{repo: r}
}
