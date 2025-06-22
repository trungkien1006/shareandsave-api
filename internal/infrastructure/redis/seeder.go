package redisapp

import (
	"context"
	"encoding/json"
	"final_project/internal/domain/redis"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/warehouse"
	"final_project/internal/pkg/enums"
	"fmt"
	"strconv"
)

type RedisSeeder struct {
	repo        redis.Repository
	rolePerRepo rolepermission.Repository
	whRepo      warehouse.Repository
}

func NewRedisSeeder(repo redis.Repository, rolePerRepo rolepermission.Repository, whRepo warehouse.Repository) *RedisSeeder {
	return &RedisSeeder{
		repo:        repo,
		rolePerRepo: rolePerRepo,
		whRepo:      whRepo,
	}
}

func (s *RedisSeeder) SeedInitialData() error {
	ctx := context.Background()

	fmt.Println("Seeding Redis initial data...")

	s.seedRolePermission(ctx)
	s.seedItemOldStock(ctx)

	fmt.Println("Seeding Redis done.")
	return nil
}

func (s *RedisSeeder) seedItemOldStock(ctx context.Context) {
	var (
		itemOldStocks []warehouse.ItemQuantity
	)

	fmt.Println("Seeding redis item old stocks...")

	itemClaimsCounts, err := s.repo.GetRedisHashLength(ctx, enums.ItemClaimRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	if itemClaimsCounts > 0 {
		fmt.Println(err)
		return
	}

	if err := s.whRepo.GetItemsQuantity(ctx, &itemOldStocks); err != nil {
		fmt.Println("Đã có dữ liệu rồi!")
		fmt.Println("Seeding redis item old stocks done.")
		return
	}

	for _, value := range itemOldStocks {
		itemClaims := warehouse.ClaimRequestItem{
			ItemQuantity: uint(value.Quantity),
			Users:        make([]warehouse.ClaimRequestUser, 0),
		}

		itemClaimsJSON, err := json.Marshal(itemClaims)
		if err != nil {
			fmt.Println("Có lỗi khi encode JSON: " + err.Error())
			return
		}

		if err := s.repo.SetToRedisHash(ctx, enums.ItemClaimRequest, "item:"+strconv.Itoa(int(value.ItemID)), string(itemClaimsJSON)); err != nil {
			fmt.Println("Có lỗi khi lưu item vào redis hash: " + err.Error())
			return
		}
	}

	fmt.Println("Seeding redis item old stocks done.")
}

func (s *RedisSeeder) seedRolePermission(ctx context.Context) {
	var (
		rolePerCodes []rolepermission.RolePermissionList
	)

	fmt.Println("Seeding redis role permissions...")

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

	fmt.Println("Seeding redis role permissions done.")
}
