package redisapp

import (
	"context"
	"final_project/internal/domain/redis"
	rolepermission "final_project/internal/domain/role_permission"
	"fmt"
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

func (s *RedisSeeder) SeedInitialData(ctx context.Context) error {
	fmt.Println("Seeding Redis initial data...")

	fmt.Println("Seeding Redis done.")
	return nil
}

func (s *RedisRepo) seedRolePermission(ctx context.Context) error {
	fmt.Println("Seeding role permissions...")

	fmt.Println("Seeding role permissions done.")
	return nil
}
