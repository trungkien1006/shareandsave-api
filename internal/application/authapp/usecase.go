package authapp

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/auth"
	"final_project/internal/domain/redis"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/user"
	"fmt"
	"strconv"
	"time"
)

type UseCase struct {
	repo         auth.Repository
	service      *auth.AuthService
	redisRepo    redis.Reposity
	roleRepo     rolepermission.Repository
	userRepo     user.Repository
	clientID     uint
	superAdminID uint
}

func NewUseCase(r auth.Repository, s *auth.AuthService, redisRepo redis.Reposity, roleRepo rolepermission.Repository, userRepo user.Repository) *UseCase {
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
	}
}

func (uc *UseCase) GetMe(ctx context.Context, user *user.User, userID uint, isAdmin bool) error {
	if isAdmin {
		if err := uc.userRepo.GetUserByID(ctx, user, int(userID), uc.clientID, uc.superAdminID); err != nil {
			return err
		}
	} else {
		if err := uc.userRepo.GetUserByID(ctx, user, int(userID), uc.clientID, 0); err != nil {
			return err
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
		return err
	}

	if isAdmin {
		// Xóa key permission:user:{userID}
		if err := uc.redisRepo.DeleteFromRedis(ctx, "permission:user:"+strconv.Itoa(int(userID))); err != nil {
			return err
		}
	}

	return nil
}
