package userapp

import (
	"context"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	"final_project/internal/dto/userDTO"
)

type UseCase struct {
	repo user.Repository
}

func NewUseCase(r user.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllUser(ctx context.Context, users *[]user.User, req userDTO.GetUserRequest) (int, error) {
	var domain_req filter.FilterRequest = filter.FilterRequest{
		Page:   req.Page,
		Limit:  req.Limit,
		Sort:   req.Sort,
		Order:  req.Order,
		Filter: req.Filter,
	}

	totalPage, err := uc.repo.GetAll(ctx, users, domain_req)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetUserByID(ctx context.Context, users *user.User, user_id int) error {
	if err := uc.repo.GetByID(ctx, users, user_id); err != nil {
		return err
	}

	return nil
}

// func (uc *UseCase) CreateUser(ctx context.Context, name, email, password string) error {

// 	exist, err := uc.repo.IsEmailExist(ctx, email)
// 	if err != nil {
// 		return err
// 	}

// 	if exist {
// 		return ErrEmailTaken
// 	}

// 	hashed, err := hash.HashPassword(password)

// 	if err != nil {
// 		return err
// 	}

// 	u := user.NewUser(name, email, hashed)

// 	return uc.repo.Save(ctx, u)
// }

// var ErrEmailTaken = errors.New("email already taken")
