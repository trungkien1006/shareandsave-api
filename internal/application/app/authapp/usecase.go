package authapp

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/auth"
	"final_project/internal/domain/email"
	"final_project/internal/domain/redis"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/user"
	authdto "final_project/internal/dto/authDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/hash"
	"final_project/internal/pkg/helpers"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type UseCase struct {
	repo         auth.Repository
	service      *auth.AuthService
	redisRepo    redis.Repository
	roleRepo     rolepermission.Repository
	userRepo     user.Repository
	clientID     uint
	superAdminID uint
	emailRepo    email.Repository
}

func NewUseCase(r auth.Repository, s *auth.AuthService, redisRepo redis.Repository, roleRepo rolepermission.Repository, userRepo user.Repository, emailRepo email.Repository) *UseCase {
	ctx := context.Background()

	clientID, err := roleRepo.GetRoleIDByName(ctx, "Client")
	if err != nil {
		fmt.Println("Có lỗi khi set roleID mặc định cho user usecase: " + err.Error())
	}

	supderAdminID, err := roleRepo.GetRoleIDByName(ctx, "Super Admin")
	if err != nil {
		fmt.Println("Có lỗi khi set roleID mặc định cho user usecase: " + err.Error())
	}

	return &UseCase{
		repo:         r,
		service:      s,
		redisRepo:    redisRepo,
		roleRepo:     roleRepo,
		userRepo:     userRepo,
		clientID:     clientID,
		superAdminID: supderAdminID,
		emailRepo:    emailRepo,
	}
}

func (uc *UseCase) VerifyOTP(ctx context.Context, req authdto.VerifyOTPRequest) (string, error) {
	//Kiểm tra số lần thử
	verifyTryCountStr, err := uc.redisRepo.GetFromRedis(ctx, "verify:email:"+req.Email+":purpose:"+req.Purpose)

	verifyTryCount := 0

	if verifyTryCountStr != "" {
		verifyTryCount, err = strconv.Atoi(verifyTryCountStr)
		if err != nil {
			return "", errors.New("Có lỗi khi chuyển đổi kiểu dữ liệu: " + err.Error())
		}
	}

	if verifyTryCount >= enums.MaxVerifyTry {
		return "", errors.New("Bạn đã thử quá 5 lần, hãy đợi 10 phút để thử lại")
	}

	//Kiểm tra OTP
	isSendedBefore, err := uc.redisRepo.GetFromRedis(ctx, "otp:email:"+req.Email+":purpose:"+req.Purpose)
	if err != nil {
		return "", errors.New("OTP không tồn tại: " + err.Error())
	}

	//Nếu OTP sai -> lưu lại lịch sử thử
	if isSendedBefore != req.OTP {
		if err := uc.redisRepo.InsertToRedis(ctx, "verify:email:"+req.Email+":purpose:"+req.Purpose, strconv.Itoa(verifyTryCount+1), 10*time.Minute); err != nil {
			return "", err
		}

		return "", errors.New("OTP không chính xác, bạn còn " + strconv.Itoa(enums.MaxVerifyTry-(verifyTryCount+1)) + " lần thử")
	}

	//Nếu đúng sẽ xóa 2 key
	if err := uc.redisRepo.DeleteFromRedis(ctx, "verify:email:"+req.Email+":purpose:"+req.Purpose); err != nil {
		return "", err
	}

	if err := uc.redisRepo.DeleteFromRedis(ctx, "otp:email:"+req.Email+":purpose:"+req.Purpose); err != nil {
		return "", err
	}

	//Tạo token xác minh các bước sau
	token, err := uc.service.GenerateVerificationToken(req.Email)
	if err != nil {
		return "", errors.New("Lỗi khi tạo token xác minh: " + err.Error())
	}

	//Lưu token vào redis
	if err := uc.redisRepo.InsertToRedis(ctx, req.Purpose+":"+token, req.Email, 10*time.Minute); err != nil {
		return "", err
	}

	return token, nil
}

func (uc *UseCase) SendOTP(ctx context.Context, req authdto.SendOTPRequest) error {
	isSendedBefore, _ := uc.redisRepo.GetFromRedis(ctx, "otp:email:"+req.Email+":purpose:"+req.Purpose)

	//Kiểm tra email tồn tại
	emailExisted, err := uc.userRepo.IsEmailExist(ctx, req.Email, 0)
	if err != nil {
		return err
	}

	if req.Purpose == "activeAccount" {
		if emailExisted {
			return errors.New("Email đã tồn tại")
		}
	} else if req.Purpose == "resetPassword" {
		if !emailExisted {
			return errors.New("Email không tồn tại")
		}
	}

	if isSendedBefore != "" {
		return errors.New("Hãy đợi đủ 5 phút để được gửi yêu cầu mới nhé")
	}

	otp := uc.service.GenerateEmailToken(req.Email)

	go func() {
		err := uc.sendOTP(otp, req.Email, "Mã xác thực của bạn")
		if err != nil {
			// return err
			fmt.Println("-----Có lỗi khi gửi mail:" + err.Error())
		}
	}()

	if err := uc.redisRepo.InsertToRedis(ctx, "otp:email:"+req.Email+":purpose:"+req.Purpose, otp, 5*time.Minute); err != nil {
		// return errors.New("Có lỗi khi lưu mã OTP: " + err.Error())
		fmt.Println("-----Có lỗi khi lưu mã OTP: " + err.Error())
	}

	return nil
}

func (uc *UseCase) VerifySignUp(ctx context.Context, req auth.AuthVerifySignUp) error {
	//Kiểm tra email tồn tại
	emailExisted, err := uc.userRepo.IsEmailExist(ctx, req.Email, 0)
	if err != nil {
		return err
	}

	if emailExisted {
		return errors.New("email: Email đã tồn tại")
	}

	//Kiểm tra số điện thoại tồn tại
	phoneNumberExisted, err := uc.userRepo.IsPhoneNumberExist(ctx, req.PhoneNumber, 0)
	if err != nil {
		return err
	}

	if phoneNumberExisted {
		return errors.New("phoneNumer: Số điện thoại đã tồn tại")
	}

	//Kiểm tra mật khẩu nhập lại chính xác
	if req.Password != req.RePassword {
		return errors.New("rePassword: Nhập lại mật khẩu không chính xác")
	}

	return nil
}

func (uc *UseCase) ClientSignUp(ctx context.Context, signUpReq auth.AuthSignUp) error {
	//Kiểm tra token hợp lệ
	isOK, err := uc.redisRepo.GetFromRedis(ctx, "activeAccount:"+signUpReq.VerifyToken)
	if err != nil {
		return errors.New("Token không tồn tại hoặc hết hạn: " + err.Error())
	}

	if isOK != signUpReq.Email {
		return errors.New("Email bạn sử dụng không phải email bạn đã xác thực")
	}

	//Kiểm tra email tồn tại
	emailExisted, err := uc.userRepo.IsEmailExist(ctx, signUpReq.Email, 0)
	if err != nil {
		return err
	}

	if emailExisted {
		return errors.New("Email đã tồn tại")
	}

	//Kiểm tra số điện thoại tồn tại
	phoneNumberExisted, err := uc.userRepo.IsPhoneNumberExist(ctx, signUpReq.PhoneNumber, 0)
	if err != nil {
		return err
	}

	if phoneNumberExisted {
		return errors.New("Số điện thoại đã tồn tại")
	}

	//Kiểm tra mật khẩu nhập lại chính xác
	if signUpReq.Password != signUpReq.RePassword {
		return errors.New("Nhập lại mật khẩu không chính xác")
	}

	var signUpUser user.User

	hashedPassword, err := hash.HashPassword(signUpReq.Password)
	if err != nil {
		return err
	}

	strBase64Image, err := helpers.ResizeImageFromFileToBase64(os.Getenv("IMAGE_PATH")+"/user.png", enums.UserImageWidth, enums.UserImageHeight)

	if err != nil {
		return err
	}

	signUpUser.RoleID = uc.clientID
	signUpUser.Email = signUpReq.Email
	signUpUser.Password = hashedPassword
	signUpUser.Avatar = strBase64Image
	signUpUser.Active = false
	signUpUser.FullName = signUpReq.FullName
	signUpUser.PhoneNumber = signUpReq.PhoneNumber
	signUpUser.Address = ""
	signUpUser.Status = int8(enums.UserStatusActive)
	signUpUser.GoodPoint = 0

	if err := uc.repo.SignUp(ctx, &signUpUser); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) ClientResetPassword(ctx context.Context, req auth.AuthResetPassword) error {
	//Kiểm tra token hợp lệ
	isOK, err := uc.redisRepo.GetFromRedis(ctx, "resetPassword:"+req.VerifyToken)
	if err != nil {
		return errors.New("Token không tồn tại hoặc hết hạn: " + err.Error())
	}

	if isOK != req.Email {
		return errors.New("Email bạn sử dụng không phải email bạn đã xác thực")
	}

	//Kiểm tra email tồn tại
	emailExisted, err := uc.userRepo.IsEmailExist(ctx, req.Email, 0)
	if err != nil {
		return err
	}

	if !emailExisted {
		return errors.New("Email không tồn tại")
	}

	if req.Password != req.RePassword {
		return errors.New("Email không tồn tại")
	}

	//Update password
	var updateUser user.User

	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return err
	}

	if err := uc.userRepo.GetByEmail(ctx, &updateUser, req.Email); err != nil {
		return err
	}

	updateUser.Password = hashedPassword

	if err := uc.userRepo.Update(ctx, &updateUser); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) sendOTP(otp, email, sub string) error {
	subject, htmlBody := uc.service.BuildOTPEmailContent(otp, sub)

	err := uc.emailRepo.Send(email, subject, htmlBody)
	if err != nil {
		return errors.New("Gửi email thất bại:" + err.Error())
	} else {
		log.Println("Đã gửi OTP thành công!")
	}

	return nil
}

func (uc *UseCase) GetMe(ctx context.Context, user *user.User, userID uint, isAdmin bool) error {
	if isAdmin {
		if err := uc.userRepo.GetMe(ctx, user, int(userID), uc.clientID, isAdmin); err != nil {
			return err
		}
	} else {
		if err := uc.userRepo.GetMe(ctx, user, int(userID), uc.clientID, isAdmin); err != nil {
			return err
		}
	}

	if isAdmin {
		// Lưu dữ liệu vào redis dưới dạng key = role:user:{userID} value = int
		if err := uc.redisRepo.InsertToRedis(ctx, "role:user:"+strconv.Itoa(int(user.ID)), string(user.RoleID), 30*24*time.Hour); err != nil {
			return err
		}

		permisisonJSON, err := uc.redisRepo.GetFromRedis(ctx, "permission:role:"+strconv.Itoa(int(user.RoleID)))
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(permisisonJSON), &user.Permissions)
		if err != nil {
			return errors.New("Có lỗi khi mã hóa danh sách quyền từ redis: " + err.Error())
		}
	}

	return nil
}

func (uc *UseCase) Login(ctx context.Context, domainAuthLogin auth.AuthLogin, JWT *string, refreshToken *string, domainUser *user.User, isAdmin bool) error {
	err := uc.repo.Login(ctx, domainUser, domainAuthLogin.Email, domainAuthLogin.Password, isAdmin, uc.clientID)
	if err != nil {
		return err
	}

	JWTSubject := auth.JWTSubject{
		Id:     domainUser.ID,
		Device: domainAuthLogin.Device,
	}

	currentVersionStr, err := uc.redisRepo.GetFromRedis(ctx, "user:"+strconv.Itoa(int(domainUser.ID))+":"+domainAuthLogin.Device)

	if err == nil && currentVersionStr != "" {
		currentVersion, err := strconv.Atoi(currentVersionStr)
		if err != nil {
			return errors.New("Có lỗi khi chuyển kiểu string sang int: " + err.Error())
		}

		JWTSubject.Version = uint(currentVersion + 1)
	} else if err != nil {
		return err
	} else {
		JWTSubject.Version = 1
	}

	*JWT = uc.service.GenerateToken(JWTSubject)
	*refreshToken = uc.service.GenerateRefreshToken(JWTSubject)

	// Lưu dữ liệu vào redis dưới dạng key = user:{userID}:{device} value = "1"
	if err := uc.redisRepo.InsertToRedis(ctx, "user:"+strconv.Itoa(int(domainUser.ID))+":"+domainAuthLogin.Device, strconv.Itoa(int(JWTSubject.Version)), 30*24*time.Hour); err != nil {
		return err
	}

	if isAdmin {
		// Lưu dữ liệu vào redis dưới dạng key = role:user:{userID} value = int
		if err := uc.redisRepo.InsertToRedis(ctx, "role:user:"+strconv.Itoa(int(domainUser.ID)), string(domainUser.RoleID), 30*24*time.Hour); err != nil {
			return err
		}

		permisisonJSON, err := uc.redisRepo.GetFromRedis(ctx, "permission:role:"+strconv.Itoa(int(domainUser.RoleID)))
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(permisisonJSON), &domainUser.Permissions)
		if err != nil {
			return errors.New("Có lỗi khi mã hóa danh sách quyền từ redis: " + err.Error())
		}
	}

	return nil
}

func (uc *UseCase) Logout(ctx context.Context, userID uint, device string, isAdmin bool) error {
	// Xóa key user:{userID}:{device}
	if err := uc.redisRepo.DeleteFromRedis(ctx, "user:"+strconv.Itoa(int(userID))+":"+device); err != nil {
		return nil
	}

	if isAdmin {
		// Xóa key permission:user:{userID}
		if err := uc.redisRepo.DeleteFromRedis(ctx, "permission:user:"+strconv.Itoa(int(userID))); err != nil {
			return nil
		}
	}

	return nil
}
