package userapp

import (
	"context"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
	"final_project/internal/pkg/helpers"
)

type UseCase struct {
	repo user.Repository
}

func NewUseCase(r user.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllUser(ctx context.Context, users *[]user.User, domain_req filter.FilterRequest) (int, error) {
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

func (uc *UseCase) CreateUser(ctx context.Context, user *user.User) error {
	emailExisted, err := uc.repo.IsEmailExist(ctx, user.Email)

	if err != nil {
		return err
	}

	if emailExisted {
		return errors.New(enums.ErrEmailExisted)
	}

	phoneNumberExisted, err := uc.repo.IsPhoneNumberExist(ctx, user.PhoneNumber)

	if err != nil {
		return err
	}

	if phoneNumberExisted {
		return errors.New(enums.ErrEmailExisted)
	}

	hashedPassword, err := hash.HashPassword(user.Password)

	if err != nil {
		return err
	}

	strBase64Image, err := helpers.ResizeImageFromFileToBase64(enums.ImagePath+"/user.jpg", enums.UserImageWidth, enums.UserImageHeight)

	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.Avatar = strBase64Image

	if err := uc.repo.Save(ctx, user); err != nil {
		return err
	}

	return nil
}
