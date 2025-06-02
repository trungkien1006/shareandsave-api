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
		Id:      domainUser.ID,
		Device:  domainAuthLogin.Device,
		Version: 1,
	}

	*JWT = uc.service.GenerateToken(JWTSubject)
	*refreshToken = uc.service.GenerateRefreshToken(JWTSubject)

	// Lưu dữ liệu vào redis dưới dạng key = user:{userID}:{device} value = "1"
	if err := uc.redisRepo.InsertToRedis(ctx, "user:"+strconv.Itoa(int(domainUser.ID))+":"+domainAuthLogin.Device, "1", 30*24*time.Hour); err != nil {
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
