package commentapp

import (
	"context"
	"final_project/internal/domain/comment"
)

type UseCase struct {
	repo comment.Repository
}

func NewUseCase(r comment.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) CreateComment(ctx context.Context, domainComment *comment.Comment) error {
	if err := uc.repo.Create(ctx, domainComment); err != nil {
		return nil
	}

	return nil
}

func (uc *UseCase) GetAllComment(ctx context.Context, domainComment *[]comment.Comment, filter comment.GetComment) error {
	if err := uc.repo.GetAll(ctx, domainComment, filter); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdateReadMessage(ctx context.Context, interestID uint) error {
	if err := uc.repo.UpdateReadMessage(ctx, interestID); err != nil {
		return err
	}

	return nil
}
