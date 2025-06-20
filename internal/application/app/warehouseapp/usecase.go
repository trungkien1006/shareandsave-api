package warehouseapp

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/item"
	"final_project/internal/domain/redis"
	"final_project/internal/domain/warehouse"
	"final_project/internal/pkg/enums"
	"strconv"
)

type UseCase struct {
	repo      warehouse.Repository
	redisRepo redis.Repository
	itemRepo  item.Repository
}

func NewUseCase(r warehouse.Repository, redisRepo redis.Repository, itemRepo item.Repository) *UseCase {
	return &UseCase{
		repo:      r,
		redisRepo: redisRepo,
		itemRepo:  itemRepo,
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
			var claimRequestItems warehouse.ClaimRequestItem

			err = json.Unmarshal([]byte(claimRequestStr), &claimRequestItems)
			if err != nil {
				return nil, 0, errors.New("Có lỗi khi giải mã JSON: " + err.Error())
			}

			claimRequestCounts = append(claimRequestCounts, uint(len(claimRequestItems.Users)))
		} else {
			claimRequestCounts = append(claimRequestCounts, 0)
		}
	}

	return claimRequestCounts, totalPage, nil
}

func (uc *UseCase) CreateClaimRequest(ctx context.Context, claimReqs []warehouse.CreateClaimRequestItem, userID uint) error {
	for _, value := range claimReqs {
		//Kiểm tra món đồ tồn tại
		itemExisted, err := uc.itemRepo.IsExist(ctx, value.ItemID)
		if err != nil {
			return err
		}

		if !itemExisted {
			return errors.New("Món đồ không tồn tại: " + err.Error())
		}

		//Lưu người dùng vào hàng đợi của từng món đồ
		itemClaimsReqJson, err := uc.redisRepo.GetFromRedisHash(ctx, enums.ItemClaimRequest, strconv.Itoa(int(value.ItemID)))
		if err != nil {
			return errors.New("Có lỗi khi truy xuất danh sách người dùng đăng kí món đồ: " + err.Error())
		}

		var itemClaims warehouse.ClaimRequestItem

		if itemClaimsReqJson != "" {
			//Decode danh sách người dùng chờ nhận đồ
			err = json.Unmarshal([]byte(itemClaimsReqJson), &itemClaims)
			if err != nil {
				return errors.New("Có lỗi khi decode JSON: " + err.Error())
			}
		}

		itemClaims.Users = append(itemClaims.Users, warehouse.ClaimRequestUser{
			ID:       userID,
			Quantity: value.Quantity,
		})

		//Kiểm tra số lượng đồ còn lại
		itemClaims.ItemQuantity += value.Quantity

		currentQuantity, err := uc.repo.GetItemWarehouseQuantity(ctx, value.ItemID)
		if err != nil {
			return err
		}

		if currentQuantity < itemClaims.ItemQuantity {
			return errors.New("Số lượng đồ đạc còn lại không đủ cho yêu cầu nhận: số lượng còn lại là " + strconv.Itoa(int(currentQuantity)))
		}

		newClaimReqJSON, err := json.Marshal(itemClaims)
		if err != nil {
			return errors.New("Có lỗi khi mã hóa JSON: " + err.Error())
		}

		if err := uc.redisRepo.SetToRedisHash(ctx, enums.ItemClaimRequest, "item:"+strconv.Itoa(int(value.ItemID)), string(newClaimReqJSON)); err != nil {
			return errors.New("Có lỗi khi lưu danh sách thành viên đang chờ nhận đồ: " + err.Error())
		}
	}

	//Lưu danh sách đăng kí nhận đồ của user vào key userClaimRequest
	claimReqsJson, err := json.Marshal(claimReqs)
	if err != nil {
		return errors.New("Có lỗi khi mã hóa JSON: " + err.Error())
	}

	if err := uc.redisRepo.SetToRedisHash(ctx, enums.UserClaimRequest, "user:"+strconv.Itoa(int(userID)), string(claimReqsJson)); err != nil {
		return errors.New("Có lỗi khi lưu danh sách đăng kí nhận đồ của người dùng: " + err.Error())
	}

	return nil
}

func (uc *UseCase) ModifyClaimRequest(ctx context.Context, domain warehouse.ModifyClaimRequest, userID uint) error {
	//Kiểm trả sản phẩm tồn tại và
	itemExisted, err := uc.itemRepo.IsExist(ctx, domain.ItemID)
	if err != nil {
		return err
	}

	if !itemExisted {
		return errors.New("Đồ đạc không tồn tại: id món đồ " + strconv.Itoa(int(domain.ItemID)))
	}

	//Chỉnh sửa số lượng trong danh sách người đăng kí trong món đồ đó
	itemClaimsReqJson, err := uc.redisRepo.GetFromRedisHash(ctx, enums.ItemClaimRequest, strconv.Itoa(int(domain.ItemID)))
	if err != nil {
		return errors.New("Có lỗi khi truy xuất danh sách người dùng đăng kí món đồ: " + err.Error())
	}

	if itemClaimsReqJson == "" {
		return errors.New("Danh sách đăng kí đồ rỗng")
	}

	var itemClaims warehouse.ClaimRequestItem

	err = json.Unmarshal([]byte(itemClaimsReqJson), &itemClaims)
	if err != nil {
		return errors.New("Có lỗi khi decode JSON: " + err.Error())
	}

	//Cập nhật số lượng món đồ ở itemClaimRequest
	oldQuantityOfItem := 0

	for key, value := range itemClaims.Users {
		if value.ID == userID {
			if domain.NewQuatity <= 0 && domain.NewQuatity > itemClaims.ItemQuantity {
				return errors.New("Số lượng đồ đạc bạn đăng kí nhận không hợp lệ: > số lượng có sẵn hoặc <= 0")
			}

			oldQuantityOfItem = int(value.Quantity)
			itemClaims.Users[key].Quantity = domain.NewQuatity
		}
	}

	itemClaims.ItemQuantity = itemClaims.ItemQuantity - uint(oldQuantityOfItem) + domain.NewQuatity

	//Cập nhật số lượng món đồ ở userClaimRequest
	userClaimsReqJson, err := uc.redisRepo.GetFromRedisHash(ctx, enums.UserClaimRequest, strconv.Itoa(int(userID)))
	if err != nil {
		return errors.New("Có lỗi khi truy xuất danh sách món đồ người dùng đăng kí: " + err.Error())
	}

	if userClaimsReqJson == "" {
		return errors.New("Danh sách món đồ của thành viên đăng kí rỗng")
	}

	var userClaims []warehouse.ClaimRequestUser

	err = json.Unmarshal([]byte(userClaimsReqJson), &userClaims)
	if err != nil {
		return errors.New("Có lỗi khi decode JSON: " + err.Error())
	}

	for key, value := range userClaims {
		if value.ID == userID {
			userClaims[key].Quantity = domain.NewQuatity
		}
	}

	return nil
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
