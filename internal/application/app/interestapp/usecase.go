package interestapp

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/domain/post"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
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

func (uc *UseCase) GetAllInterest(ctx context.Context, postInterest *[]interest.PostInterest, userID uint, filter interest.GetInterest) (uint, int, error) {
	isUserExist, err := uc.userRepo.IsExist(ctx, userID)
	if err != nil {
		return 0, 0, err
	}

	if !isUserExist {
		return 0, 0, errors.New("Người dùng không tồn tại")
	}

	totalPage, err := uc.repo.GetAll(ctx, postInterest, userID, filter)

	if err != nil {
		return 0, 0, err
	}

	unreadMessageCount, err := uc.repo.GetTotalUnreadMessage(ctx, userID, enums.InterestType(filter.Type))
	if err != nil {
		return 0, 0, err
	}

	return unreadMessageCount, totalPage, nil
}

func (uc *UseCase) GetInterestByID(ctx context.Context, postInterest *interest.PostInterest, interestID uint) error {
	err := uc.repo.GetDetailByID(ctx, postInterest, interestID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) CreateInterest(ctx context.Context, interest interest.Interest) (uint, error) {
	isUserExist, err := uc.userRepo.IsExist(ctx, interest.UserID)
	if err != nil {
		return 0, err
	}

	if !isUserExist {
		return 0, errors.New("Người dùng không tồn tại")
	}

	isPostExist, err := uc.postRepo.IsExist(ctx, interest.PostID)
	if err != nil {
		return 0, err
	}

	if !isPostExist {
		return 0, errors.New("Bài viết không tồn tại")
	}

	isInterestExist, err := uc.repo.IsExist(ctx, interest.UserID, interest.PostID)
	if err != nil {
		return 0, err
	}

	if isInterestExist {
		return 0, errors.New("Bạn đã quan tâm rồi")
	}

	id, err := uc.repo.Create(ctx, interest)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *UseCase) DeleteInterest(ctx context.Context, postID uint, userID uint) (uint, error) {
	id, err := uc.repo.Delete(ctx, postID, userID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
