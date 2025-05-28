package requestapp

import (
	"context"
	"errors"
	"final_project/internal/domain/request"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
	"final_project/internal/pkg/helpers"
	"fmt"
	"os"
)

type UseCase struct {
	repo     request.Repository
	userRepo user.Repository
}

func NewUseCase(r request.Repository, userRepo user.Repository) *UseCase {
	return &UseCase{
		repo:     r,
		userRepo: userRepo,
	}
}

func (uc *UseCase) CreateRequest(ctx context.Context, req *request.Request, user *user.User) error {
	//Nếu không truyền userID sẽ kiểm tra để tạo tài khoản cho người dùng
	if req.UserID == 0 {
		//Kiểm tra email đã tồn tại trong hệ thống chưa
		userEmailExist, err := uc.userRepo.IsEmailExist(ctx, user.Email)
		if err != nil {
			return err
		}

		//Kiểm tra sđt đã tồn tại trong hệ thống chưa
		userPhoneNumberExist, err := uc.userRepo.IsPhoneNumberExist(ctx, user.PhoneNumber)
		if err != nil {
			return err
		}

		//Nếu đã tồn tại cả email và số điện thoại thì lấy ID của người dùng, không thì trả lỗi nếu 1 trong 2 tồn tại rồi
		if userEmailExist && userPhoneNumberExist {
			req.UserID, err = uc.userRepo.GetIDByEmailPhoneNumber(ctx, user.Email, user.PhoneNumber)

			if err != nil {
				return err
			}
		} else if userEmailExist {
			return errors.New("Đã có tài khoản sở hữu email này")
		} else if userPhoneNumberExist {
			return errors.New("Đã có tài khoản sở hữu số điện thoại này")
		} else {
			//Nếu chưa tồn tại cả email và số điện thoại thì tạo mới người dùng

			strBase64Image, err := helpers.ResizeImageFromFileToBase64(os.Getenv("IMAGE_PATH")+"/user.png", enums.UserImageWidth, enums.UserImageHeight)
			if err != nil {
				return fmt.Errorf("error resizing image: %w", err)
			}

			user.Avatar = strBase64Image
			user.Password = hash.HashEmailPhone(user.Email, user.PhoneNumber) // Mã hóa mật khẩu bằng email và số điện thoại
			user.Address = ""
			user.Status = int(enums.UserStatusInactive)
			user.GoodPoint = 0
			user.ID = 0 // Đặt ID về 0 để đảm bảo tạo mới

			if err := uc.userRepo.Save(ctx, user); err != nil {
				return err
			}

			req.UserID = user.ID
		}
	}

	req.Status = int8(enums.RequestStatusPending) // Mặc định trạng thái là Pending
	req.ReplyMessage = ""                         // Mặc định ReplyMessage là rỗng

	if err := uc.repo.Create(ctx, req); err != nil {
		return err
	}

	return nil
}
