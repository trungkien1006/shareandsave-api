package interestapp

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/domain/post"
	"final_project/internal/domain/user"
)

type UseCase struct {
	repo     interest.Repository
	userRepo user.Repository
	postRepo post.Repository
}

func NewUseCase(r interest.Repository, userRepo user.Repository, postRepo post.Repository) *UseCase {
	return &UseCase{
		repo:     r,
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (uc *UseCase) GetAllInterest(ctx context.Context, postInterest *[]interest.PostInterest, userID uint, filter interest.GetInterest) (int, error) {
	isUserExist, err := uc.userRepo.IsExist(ctx, userID)
	if err != nil {
		return 0, err
	}

	if !isUserExist {
		return 0, errors.New("Người dùng không tồn tại")
	}

	totalPage, err := uc.repo.GetAll(ctx, postInterest, userID, filter)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) CreateInterest(ctx context.Context, interest interest.Interest) error {
	isUserExist, err := uc.userRepo.IsExist(ctx, interest.UserID)
	if err != nil {
		return err
	}

	if !isUserExist {
		return errors.New("Người dùng không tồn tại")
	}

	isPostExist, err := uc.postRepo.IsExist(ctx, interest.PostID)
	if err != nil {
		return err
	}

	if !isPostExist {
		return errors.New("Bài viết không tồn tại")
	}

	isInterestExist, err := uc.repo.IsExist(ctx, interest.UserID, interest.PostID)
	if err != nil {
		return err
	}

	if isInterestExist {
		return errors.New("Bạn đã quan tâm rồi")
	}

	if err := uc.repo.Create(ctx, interest); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) DeleteInterest(ctx context.Context, interestID uint, userID uint) error {
	if err := uc.repo.Delete(ctx, interestID); err != nil {
		return err
	}

	return nil
}
