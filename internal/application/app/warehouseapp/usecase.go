package warehouseapp

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/redis"
	"final_project/internal/domain/warehouse"
	"final_project/internal/pkg/enums"
	"strconv"
)

type UseCase struct {
	repo      warehouse.Repository
	redisRepo redis.Repository
}

func NewUseCase(r warehouse.Repository, redisRepo redis.Repository) *UseCase {
	return &UseCase{
		repo:      r,
		redisRepo: redisRepo,
	}
}

func (uc *UseCase) GetAllWarehouse(ctx context.Context, warehouses *[]warehouse.Warehouse, filter filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, warehouses, filter)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetAllItemWarehouse(ctx context.Context, warehouses *[]warehouse.ItemWareHouse, filter filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAllItem(ctx, warehouses, filter)
	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetAllItemOldStock(ctx context.Context, items *[]warehouse.ItemOldStock, filter filter.FilterRequest) ([]uint, int, error) {
	totalPage, err := uc.repo.GetAllOldStockItem(ctx, items, filter)
	if err != nil {
		return nil, 0, err
	}

	var claimRequestCounts []uint

	for _, value := range *items {
		claimRequestStr, err := uc.redisRepo.GetFromRedisHash(ctx, enums.ItemClaimRequest, strconv.Itoa(int(value.ItemID)))
		if err != nil {
			return nil, 0, errors.New("Có lỗi khi truy xuất số thành viên đã đăng kí nhận đồ: " + err.Error())
		}

		if claimRequestStr != "" {
			var claimRequests []warehouse.ClaimRequest

			err = json.Unmarshal([]byte(claimRequestStr), &claimRequests)
			if err != nil {
				return nil, 0, errors.New("Có lỗi khi giải mã JSON: " + err.Error())
			}

			claimRequestCounts = append(claimRequestCounts, uint(len(claimRequests)))
		} else {
			claimRequestCounts = append(claimRequestCounts, 0)
		}
	}

	return claimRequestCounts, totalPage, nil
}

func (uc *UseCase) GetItemByCode(ctx context.Context, itemWarehouse *warehouse.ItemWareHouse, code string) error {
	if err := uc.repo.GetItemByCode(ctx, itemWarehouse, code); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetWarehouseByID(ctx context.Context, warehouse *warehouse.DetailWarehouse, warehouseID uint) error {
	if err := uc.repo.GetByID(ctx, warehouse, warehouseID); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdateWarehouse(ctx context.Context, domainWarehouse warehouse.DetailWarehouse) error {
	var updateWarehouse warehouse.DetailWarehouse

	if domainWarehouse.Description != "" {
		updateWarehouse.Description = domainWarehouse.Description
	}

	if domainWarehouse.StockPlace != "" {
		updateWarehouse.StockPlace = domainWarehouse.StockPlace
	}

	if domainWarehouse.ItemWareHouse != nil {
		updateWarehouse.ItemWareHouse = domainWarehouse.ItemWareHouse
	}

	updateWarehouse.ID = domainWarehouse.ID

	if err := uc.repo.Update(ctx, updateWarehouse); err != nil {
		return err
	}

	return nil
}
