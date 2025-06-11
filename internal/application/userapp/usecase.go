package userapp

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/redis"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
	"final_project/internal/pkg/helpers"
	"fmt"
	"os"
	"strconv"
	"time"
)

type UseCase struct {
	repo         user.Repository
	roleRepo     rolepermission.Repository
	clientID     uint
	superAdminID uint
	redisRepo    redis.Repository
}

func NewUseCase(r user.Repository, roleRepo rolepermission.Repository, redisRepo redis.Repository) *UseCase {
	ctx := context.Background()

	clientID, err := roleRepo.GetRoleIDByName(ctx, "Client")
	if err != nil {
		fmt.Println("Có lỗi khi set clientID cho user usecase: " + err.Error())
	}

	supderAdminID, err := roleRepo.GetRoleIDByName(ctx, "Super Admin")
	if err != nil {
		fmt.Println("Có lỗi khi set superAdminID cho user usecase: " + err.Error())
	}

	return &UseCase{
		repo:         r,
		roleRepo:     roleRepo,
		clientID:     uint(clientID),
		superAdminID: supderAdminID,
		redisRepo:    redisRepo,
	}
}

func (uc *UseCase) GetAllClient(ctx context.Context, users *[]user.User, domainReq filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, users, domainReq, uc.clientID, 0)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetAllAdmin(ctx context.Context, users *[]user.User, domainReq filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, users, domainReq, uc.clientID, uc.superAdminID)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetClientByID(ctx context.Context, users *user.User, userID int) error {
	if err := uc.repo.GetUserByID(ctx, users, userID, uc.clientID, 0); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetAdminByID(ctx context.Context, users *user.User, userID int) error {
	if err := uc.repo.GetUserByID(ctx, users, userID, uc.clientID, uc.superAdminID); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) CreateClient(ctx context.Context, user *user.User) error {
	roleExisted, err := uc.roleRepo.IsRoleExisted(ctx, uc.clientID)

	if err != nil {
		return err
	}

	if !roleExisted {
		return errors.New(enums.ErrRoleNotExist)
	}

	emailExisted, err := uc.repo.IsEmailExist(ctx, user.Email, 0)

	if err != nil {
		return err
	}

	if emailExisted {
		return errors.New(enums.ErrEmailExisted)
	}

	phoneNumberExisted, err := uc.repo.IsPhoneNumberExist(ctx, user.PhoneNumber, 0)

	if err != nil {
		return err
	}

	if phoneNumberExisted {
		return errors.New(enums.ErrPhoneNumberExisted)
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
	user.RoleID = uc.clientID

	if err := uc.repo.Save(ctx, user); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) CreateAdmin(ctx context.Context, user *user.User) error {
	roleExisted, err := uc.roleRepo.IsRoleExisted(ctx, user.RoleID)

	if err != nil {
		return err
	}

	if !roleExisted {
		return errors.New(enums.ErrRoleNotExist)
	}

	emailExisted, err := uc.repo.IsEmailExist(ctx, user.Email, 0)

	if err != nil {
		return err
	}

	if emailExisted {
		return errors.New(enums.ErrEmailExisted)
	}

	phoneNumberExisted, err := uc.repo.IsPhoneNumberExist(ctx, user.PhoneNumber, 0)

	if err != nil {
		return err
	}

	if phoneNumberExisted {
		return errors.New(enums.ErrPhoneNumberExisted)
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
	user.RoleID = uc.clientID

	if err := uc.repo.Save(ctx, user); err != nil {
		return err
	}

	// Lưu dữ liệu vào redis dưới dạng key = role:user:{userID} value = role_id
	if err := uc.redisRepo.InsertToRedis(ctx, "role:user:"+strconv.Itoa(int(user.ID)), string(user.RoleID), 30*24*time.Hour); err != nil {
		return err
	}

	permissionJSON, err := uc.redisRepo.GetFromRedis(ctx, "permission:role"+strconv.Itoa(int(user.ID)))
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(permissionJSON), &user.Permissions)
	if err != nil {
		return errors.New("Có lỗi khi mã hóa danh sách quyền: " + err.Error())
	}

	return nil
}

func (uc *UseCase) UpdateClient(ctx context.Context, domainUser *user.User) error {
	var updateUser user.User

	if err := uc.repo.GetUserByID(ctx, &updateUser, int(domainUser.ID), uc.clientID, 0); err != nil {
		return err
	}

	if updateUser.ID == 0 {
		return errors.New(enums.ErrUserNotExist)
	}

	if domainUser.PhoneNumber != "" {
		phoneNumberExisted, err := uc.repo.IsPhoneNumberExist(ctx, domainUser.PhoneNumber, int(domainUser.ID))

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

	if domainUser.Major != "" {
		updateUser.Major = domainUser.Major
	}

	updateUser.Status = domainUser.Status
	updateUser.RoleID = uc.clientID

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

func (uc *UseCase) UpdateAdmin(ctx context.Context, domainUser *user.User) error {
	var updateUser user.User

	if err := uc.repo.GetUserByID(ctx, &updateUser, int(domainUser.ID), uc.clientID, uc.superAdminID); err != nil {
		return err
	}

	if updateUser.ID == 0 {
		return errors.New(enums.ErrUserNotExist)
	}

	if domainUser.PhoneNumber != "" {
		phoneNumberExisted, err := uc.repo.IsPhoneNumberExist(ctx, domainUser.PhoneNumber, int(domainUser.ID))

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

	if domainUser.Major != "" {
		updateUser.Major = domainUser.Major
	}

	updateUser.Status = domainUser.Status
	updateUser.RoleID = domainUser.RoleID

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

	// Lưu dữ liệu vào redis dưới dạng key = role:user:{userID} value = role_id
	if err := uc.redisRepo.InsertToRedis(ctx, "role:user:"+strconv.Itoa(int(domainUser.ID)), string(domainUser.RoleID), 30*24*time.Hour); err != nil {
		return err
	}

	permissionJSON, err := uc.redisRepo.GetFromRedis(ctx, "permission:role"+strconv.Itoa(int(domainUser.ID)))
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(permissionJSON), &domainUser.Permissions)
	if err != nil {
		return errors.New("Có lỗi khi mã hóa danh sách quyền: " + err.Error())
	}

	return nil
}

func (uc *UseCase) DeleteClient(ctx context.Context, userID int) error {
	var deleteUser user.User

	if err := uc.repo.GetUserByID(ctx, &deleteUser, int(userID), uc.clientID, 0); err != nil {
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

func (uc *UseCase) DeleteAdmin(ctx context.Context, userID int) error {
	var deleteUser user.User

	if err := uc.repo.GetUserByID(ctx, &deleteUser, int(userID), uc.clientID, uc.superAdminID); err != nil {
		return err
	}

	if deleteUser.ID == 0 {
		return errors.New(enums.ErrUserNotExist)
	}

	if err := uc.repo.Delete(ctx, &deleteUser); err != nil {
		return err
	}

	if err := uc.redisRepo.DeleteFromRedis(ctx, "role:user:"+strconv.Itoa(int(deleteUser.ID))); err != nil {
		return errors.New("Có lỗi khi xóa role id của user: " + err.Error())
	}

	return nil
}
