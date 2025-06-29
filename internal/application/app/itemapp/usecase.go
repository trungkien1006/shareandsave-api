package itemapp

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/item"
	"final_project/internal/domain/redis"
	"final_project/internal/domain/warehouse"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"os"
	"strconv"
)

type UseCase struct {
	repo      item.Repository
	redisRepo redis.Repository
}

func NewUseCase(r item.Repository, redisRepo redis.Repository) *UseCase {
	return &UseCase{repo: r, redisRepo: redisRepo}
}

func (uc *UseCase) GetAllItem(ctx context.Context, items *[]item.Item, domainReq filter.FilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, items, domainReq)
	if err != nil {
		return 0, err
	}
	return totalPage, nil
}

func (uc *UseCase) GetItemByID(ctx context.Context, entity *item.Item, itemID uint) error {
	if err := uc.repo.GetByID(ctx, entity, itemID); err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) CreateItem(ctx context.Context, entity *item.Item) error {
	// Xử lý ảnh giống userUsecase
	if entity.Image == "" {
		strBase64Image, err := helpers.ResizeImageFromFileToBase64(os.Getenv("IMAGE_PATH")+"/item.png", enums.ItemImageWidth, enums.ItemImageHeight)
		if err != nil {
			return err
		}

		entity.Image = strBase64Image
	}

	if err := uc.repo.Save(ctx, entity); err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) UpdateItem(ctx context.Context, domainItem *item.Item) error {
	var updateItem item.Item

	if err := uc.repo.GetByID(ctx, &updateItem, domainItem.ID); err != nil {
		return err
	}

	if updateItem.ID == 0 {
		return errors.New(enums.ErrItemNotExist)
	}

	if domainItem.Name != "" {
		updateItem.Name = domainItem.Name
	}

	if domainItem.Description != "" {
		updateItem.Description = domainItem.Description
	}

	if domainItem.CategoryID != 0 {
		updateItem.CategoryID = domainItem.CategoryID
	}

	if domainItem.MaxClaim != 0 {
		updateItem.MaxClaim = domainItem.MaxClaim

		itemClaimJSON, err := uc.redisRepo.GetFromRedisHash(ctx, enums.ItemClaimRequest, "item:"+strconv.Itoa(int(domainItem.ID)))
		if err != nil {
			return errors.New("Có lỗi khi lấy dữ liệu từ Redis: " + err.Error())
		}

		if itemClaimJSON != "" {
			var itemClaims warehouse.ClaimRequestItem

			err = json.Unmarshal([]byte(itemClaimJSON), &itemClaims)
			if err != nil {
				return errors.New("Có lỗi khi decode JSON error: " + err.Error())
			}

			itemClaims.MaxClaim = uint(domainItem.MaxClaim)

			newItemClaimJSON, err := json.Marshal(itemClaims)
			if err != nil {
				return errors.New("Có lỗi khi encode JSON error: " + err.Error())
			}

			if err := uc.redisRepo.SetToRedisHash(ctx, enums.ItemClaimRequest, "item:"+strconv.Itoa(int(domainItem.ID)), string(newItemClaimJSON)); err != nil {
				return err
			}
		}
	}

	updateItem.ID = domainItem.ID

	if domainItem.Image != "" {
		strBase64Image, err := helpers.ProcessImageBase64(domainItem.Image, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)
		if err != nil {
			return err
		}
		updateItem.Image = strBase64Image
	}

	if err := uc.repo.Update(ctx, &updateItem); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) DeleteItem(ctx context.Context, itemID uint) error {
	var deleteItem item.Item

	if err := uc.repo.GetByID(ctx, &deleteItem, itemID); err != nil {
		return err
	}

	if deleteItem.ID == 0 {
		return errors.New(enums.ErrItemNotExist)
	}

	if err := uc.repo.Delete(ctx, &deleteItem); err != nil {
		return err
	}

	return nil
}
