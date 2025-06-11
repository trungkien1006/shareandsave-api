package redisapp

import (
	"context"
	"encoding/json"
	"final_project/internal/domain/redis"
	rolepermission "final_project/internal/domain/role_permission"
	"fmt"
	"strconv"
)

type RedisSeeder struct {
	repo        redis.Repository
	rolePerRepo rolepermission.Repository
}

func NewRedisSeeder(repo redis.Repository, rolePerRepo rolepermission.Repository) *RedisSeeder {
	return &RedisSeeder{
		repo:        repo,
		rolePerRepo: rolePerRepo,
	}
}

func (s *RedisSeeder) SeedInitialData() error {
	ctx := context.Background()

	fmt.Println("Seeding Redis initial data...")

	s.seedRolePermission(ctx)

	fmt.Println("Seeding Redis done.")
	return nil
}

func (s *RedisSeeder) seedRolePermission(ctx context.Context) {
	var (
		rolePerCodes []rolepermission.RolePermissionList
	)

	fmt.Println("Seeding role permissions...")

	if err := s.rolePerRepo.GetAllRolePermisson(ctx, &rolePerCodes); err != nil {
		fmt.Println("Có lỗi khi lấy danh sách quyền theo role")
		return
	}

	for _, value := range rolePerCodes {
		permissionJSON, err := json.Marshal(value.Permissions)
		if err != nil {
			fmt.Println("Có lỗi khi mã hóa quyền thành JSON")
			return
		}

		if err := s.repo.InsertToRedis(ctx, "permission:role:"+strconv.Itoa(int(value.ID)), string(permissionJSON), 0); err != nil {
			fmt.Println("Có lỗi khi lưu danh sách quyền vào redis")
			return
		}
	}

	fmt.Println("Seeding role permissions done.")
}
