package authapp

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/auth"
	"final_project/internal/domain/redis"
	"final_project/internal/domain/user"
	"strconv"
	"time"
)

type UseCase struct {
	repo      auth.Repository
	service   *auth.AuthService
	redisRepo redis.Reposity
}

func NewUseCase(r auth.Repository, s *auth.AuthService, redisRepo redis.Reposity) *UseCase {
	return &UseCase{
		repo:      r,
		service:   s,
		redisRepo: redisRepo,
	}
}

func (uc *UseCase) Login(ctx context.Context, domainAuthLogin auth.AuthLogin, JWT *string, refreshToken *string, domainUser *user.User) error {
	permissionCodes, err := uc.repo.Login(ctx, domainUser, domainAuthLogin.Email, domainAuthLogin.Password)
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

	if permissionCodes != nil {
		permissionCodesJSON, err := json.Marshal(permissionCodes)
		if err != nil {
			return errors.New("Lỗi khi mã hóa danh sách quyền của user: " + err.Error())
		}

		// Lưu dữ liệu vào redis dưới dạng key = permission:user:{userID} value = "[]string"
		if err := uc.redisRepo.InsertToRedis(ctx, "permission:user:"+strconv.Itoa(int(domainUser.ID)), string(permissionCodesJSON), 30*24*time.Hour); err != nil {
			return err
		}
	}

	return nil
}

func (uc *UseCase) Logout(ctx context.Context, userID uint, device string) error {
	// Xóa key user:{userID}:{device}
	if err := uc.redisRepo.DeleteFromRedis(ctx, "user:"+strconv.Itoa(int(userID))+":"+device); err != nil {
		return err
	}

	// Xóa key permission:user:{userID}
	if err := uc.redisRepo.DeleteFromRedis(ctx, "permission:user:"+strconv.Itoa(int(userID))); err != nil {
		return err
	}

	return nil
}
