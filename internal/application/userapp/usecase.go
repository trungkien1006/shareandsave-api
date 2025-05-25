package userapp

import (
	"context"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
	"final_project/internal/pkg/helpers"
	"os"
)

type UseCase struct {
	repo user.Repository
}

func NewUseCase(r user.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllUser(ctx context.Context, users *[]user.User, domainReq filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, users, domainReq)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetUserByID(ctx context.Context, users *user.User, userID int) error {
	if err := uc.repo.GetByID(ctx, users, userID); err != nil {
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

	strBase64Image, err := helpers.ResizeImageFromFileToBase64(os.Getenv("IMAGE_PATH")+"/user.png", enums.UserImageWidth, enums.UserImageHeight)

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

func (uc *UseCase) UpdateUser(ctx context.Context, domainUser *user.User) error {
	var updateUser user.User

	if err := uc.repo.GetByID(ctx, &updateUser, int(domainUser.ID)); err != nil {
		return err
	}

	if updateUser.ID == 0 {
		return errors.New(enums.ErrUserNotExist)
	}

	if domainUser.Email != "" {
		emailExisted, err := uc.repo.IsEmailExist(ctx, domainUser.Email)

		if err != nil {
			return err
		}

		if emailExisted {
			return errors.New(enums.ErrEmailExisted)
		}

		updateUser.Email = domainUser.Email
	}

	if domainUser.PhoneNumber != "" {
		phoneNumberExisted, err := uc.repo.IsPhoneNumberExist(ctx, domainUser.PhoneNumber)

		if err != nil {
			return err
		}

		if phoneNumberExisted {
			return errors.New(enums.ErrPhoneNumberExisted)
		}

		updateUser.PhoneNumber = domainUser.PhoneNumber
	}

	if domainUser.FullName != "" {
		updateUser.FullName = domainUser.FullName
	}

	if domainUser.Address != "" {
		updateUser.Address = domainUser.Address
	}

	if domainUser.GoodPoint >= 0 {
		updateUser.GoodPoint = domainUser.GoodPoint
	}

	updateUser.Status = domainUser.Status

	if domainUser.Password != "" {
		hashedPassword, err := hash.HashPassword(domainUser.Password)

		if err != nil {
			return err
		}

		updateUser.Password = hashedPassword
	}

	if domainUser.Avatar != "" {
		strBase64Image, err := helpers.ResizeImageFromBase64(domainUser.Avatar, enums.UserImageWidth, enums.UserImageHeight)

		if err != nil {
			return err
		}

		updateUser.Avatar = strBase64Image
	}

	if err := uc.repo.Update(ctx, &updateUser); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) DeleteUser(ctx context.Context, userID int) error {
	var deleteUser user.User

	if err := uc.repo.GetByID(ctx, &deleteUser, int(userID)); err != nil {
		return err
	}

	if deleteUser.ID == 0 {
		return errors.New(enums.ErrUserNotExist)
	}

	if err := uc.repo.Delete(ctx, &deleteUser); err != nil {
		return err
	}

	return nil
}
