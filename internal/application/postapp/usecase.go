package postapp

import (
	"context"
	"errors"
	"final_project/internal/domain/post"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
	"final_project/internal/pkg/helpers"
	"fmt"
	"os"
)

type UseCase struct {
	repo     post.Repository
	service  post.PostService
	userRepo user.Repository
	roleRepo rolepermission.Repository
}

func NewUseCase(r post.Repository, userRepo user.Repository, roleRepo rolepermission.Repository) *UseCase {
	return &UseCase{
		repo:     r,
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (uc *UseCase) CreatePost(ctx context.Context, post *post.Post, user *user.User) (string, error) {
	//Nếu không truyền userID sẽ kiểm tra để tạo tài khoản cho người dùng
	if post.AuthorID == 0 {
		//Kiểm tra email đã tồn tại trong hệ thống chưa
		userEmailExist, err := uc.userRepo.IsEmailExist(ctx, user.Email, 0)
		if err != nil {
			return "", err
		}

		//Kiểm tra sđt đã tồn tại trong hệ thống chưa
		userPhoneNumberExist, err := uc.userRepo.IsPhoneNumberExist(ctx, user.PhoneNumber, 0)
		if err != nil {
			return "", err
		}

		//Nếu đã tồn tại cả email và số điện thoại thì lấy ID của người dùng, không thì trả lỗi nếu 1 trong 2 tồn tại rồi
		if userEmailExist && userPhoneNumberExist {
			post.AuthorID, err = uc.userRepo.GetIDByEmailPhoneNumber(ctx, user.Email, user.PhoneNumber)

			if err != nil {
				return "", err
			}
		} else if userEmailExist {
			return "", errors.New("Đã có tài khoản sở hữu email này")
		} else if userPhoneNumberExist {
			return "", errors.New("Đã có tài khoản sở hữu số điện thoại này")
		} else {
			//Nếu chưa tồn tại cả email và số điện thoại thì tạo mới người dùng

			strBase64Image, err := helpers.ResizeImageFromFileToBase64(os.Getenv("IMAGE_PATH")+"/user.png", enums.UserImageWidth, enums.UserImageHeight)
			if err != nil {
				return "", fmt.Errorf("Lỗi khi resize ảnh: %w", err)
			}

			clientRoleID, err := uc.roleRepo.GetRoleIDByName(ctx, "Client")
			if err != nil {
				return "", errors.New("Lỗi khi lấy ID chức vụ của client:" + err.Error())
			}

			user.RoleID = clientRoleID
			user.Avatar = strBase64Image
			user.Password = hash.HashEmailPhone(user.Email, user.PhoneNumber) // Mã hóa mật khẩu bằng email và số điện thoại
			user.Address = ""
			user.Status = int8(enums.UserStatusInactive)
			user.GoodPoint = 0
			user.ID = 0 // Đặt ID về 0 để đảm bảo tạo mới

			if err := uc.userRepo.Save(ctx, user); err != nil {
				return "", err
			}

			post.AuthorID = user.ID
		}
	}

	if post.AuthorID == 0 {
		return "", errors.New("Không thể tạo bài viết mà không có người dùng")
	}

	post.Status = int8(enums.PostStatusPending) // Mặc định trạng thái là Pending
	post.Content = uc.service.GenerateContent(post.Info)
	post.Slug = uc.service.GenerateSlug(post.Title)
	post.AuthorName = user.FullName

	//resize ảnh
	for index, image := range post.Images {
		formatedImage, err := helpers.ProcessImageBase64(image, uint(enums.PostImageWidth), uint(enums.PostImageHeight), 75, helpers.FormatJPEG)
		if err != nil {
			return "", errors.New("Không thể format ảnh:" + err.Error())
		}

		post.Images[index] = formatedImage
	}

	//create post
	if err := uc.repo.Save(ctx, post); err != nil {
		return "", err
	}

	userJWTSub := helpers.UserJWTSubject{
		Id:   user.ID,
		Name: user.FullName,
	}

	JWT := helpers.GenerateToken(userJWTSub)

	return JWT, nil
}
