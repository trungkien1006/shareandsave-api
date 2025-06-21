package postapp

import (
	"context"
	"errors"
	"final_project/internal/domain/category"
	"final_project/internal/domain/item"
	"final_project/internal/domain/post"
	rolepermission "final_project/internal/domain/role_permission"
	"final_project/internal/domain/user"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"os"
	"time"
)

type UseCase struct {
	repo         post.Repository
	service      *post.PostService
	userRepo     user.Repository
	roleRepo     rolepermission.Repository
	itemRepo     item.Repository
	categoryRepo category.Repository
}

func NewUseCase(r post.Repository, userRepo user.Repository, roleRepo rolepermission.Repository, service *post.PostService, itemRepo item.Repository, categoryRepo category.Repository) *UseCase {
	return &UseCase{
		repo:         r,
		service:      service,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		itemRepo:     itemRepo,
		categoryRepo: categoryRepo,
	}
}

func (uc *UseCase) GetAllAdminPost(ctx context.Context, posts *[]post.Post, filter post.AdminPostFilterRequest, userID uint) (int, error) {
	totalPage, err := uc.repo.AdminGetAll(ctx, posts, filter, userID)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetAllPost(ctx context.Context, posts *[]post.PostWithCount, filter post.PostFilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAll(ctx, posts, filter)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetAllUserPost(ctx context.Context, posts *[]post.PostWithCount, userID uint, filter post.PostFilterRequest) (int, error) {
	totalPage, err := uc.repo.GetAllUserPost(ctx, posts, userID, filter)

	if err != nil {
		return 0, err
	}

	return totalPage, nil
}

func (uc *UseCase) GetPostByID(ctx context.Context, post *post.DetailPost, postID uint) error {
	if err := uc.repo.GetDetailByID(ctx, post, postID); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetPostBySlug(ctx context.Context, post *post.DetailPost, postSlug string) error {
	if err := uc.repo.GetDetailBySlug(ctx, post, postSlug); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) CreatePost(ctx context.Context, domainPost *post.CreatePost) error {
	var (
		author user.User
		tags   []string
	)

	itemIDs := make(map[uint]uint, 0)

	if err := uc.userRepo.GetCommonUserByID(ctx, &author, int(domainPost.AuthorID)); err != nil {
		return err
	}

	if domainPost.Info != "" {
		postContent, err := uc.service.GenerateContent(domainPost.Info)
		if err != nil {
			return errors.New("Lỗi khi tạo content từ info:" + err.Error())
		}

		domainPost.Content = postContent
	} else {
		domainPost.Content = "{}"
		domainPost.Info = "{}"
	}

	domainPost.AuthorID = author.ID
	domainPost.Status = int8(enums.PostStatusPending) // Mặc định trạng thái là Pending
	domainPost.Slug = uc.service.GenerateSlug(domainPost.Title)
	domainPost.AuthorName = author.FullName

	//resize ảnh
	for index, image := range domainPost.Images {
		formatedImage, err := helpers.ProcessImageBase64(image, uint(enums.PostImageWidth), uint(enums.PostImageHeight), 75, helpers.FormatJPEG)
		if err != nil {
			return errors.New("Không thể format ảnh:" + err.Error())
		}

		domainPost.Images[index] = formatedImage
	}

	for key, oldItem := range domainPost.OldItems {
		var item item.Item

		err := uc.itemRepo.GetByID(ctx, &item, oldItem.ItemID)
		if err != nil {
			return err
		}

		if oldItem.Image == "" {
			domainPost.OldItems[key].Image = item.Image
		} else {
			strBase64Image, err := helpers.ProcessImageBase64(oldItem.Image, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)

			if err != nil {
				return err
			}

			domainPost.OldItems[key].Image = strBase64Image
		}

		itemIDs[oldItem.ItemID] = oldItem.ItemID
	}

	for key, newItem := range domainPost.NewItems {
		var checkItem item.Item

		err := uc.itemRepo.GetByName(ctx, &checkItem, newItem.Name)

		if err != nil {
			if newItem.Image == "" {
				base64, err := helpers.ImageToBase64(os.Getenv("IMAGE_PATH") + "/item.png")
				if err != nil {
					return err
				}

				strBase64Image, err := helpers.ProcessImageBase64(base64, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)
				if err != nil {
					return err
				}

				checkItem.Image = strBase64Image
			} else {
				strBase64Image, err := helpers.ProcessImageBase64(newItem.Image, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)

				if err != nil {
					return err
				}

				checkItem.Image = strBase64Image
			}

			checkItem.CategoryID = newItem.CategoryID
			checkItem.Name = newItem.Name
			checkItem.Description = ""

			if err := uc.itemRepo.Save(ctx, &checkItem); err != nil {
				return err
			}

			domainPost.NewItems[key].ItemID = checkItem.ID
		} else {
			var tempItem post.OldItemsInPost

			if newItem.Image != "" {
				strBase64Image, err := helpers.ProcessImageBase64(newItem.Image, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)
				if err != nil {
					return err
				}

				tempItem.Image = strBase64Image
			} else {
				tempItem.Image = newItem.Image
			}

			tempItem.ItemID = checkItem.ID
			tempItem.Quantity = newItem.Quantity
			tempItem.CurrentQuantity = newItem.CurrentQuantity

			domainPost.OldItems = append(domainPost.OldItems, tempItem)
		}

		itemIDs[checkItem.ID] = checkItem.ID
	}

	if err := uc.categoryRepo.GetCategoryNameByItemIDs(ctx, itemIDs, &tags); err != nil {
		return err
	}

	domainPost.Tag = tags

	//create post
	if err := uc.repo.Save(ctx, domainPost); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdatePost(ctx context.Context, domainPost *post.Post, isRepost bool) error {
	var updatePost post.Post

	if err := uc.repo.GetByID(ctx, &updatePost, domainPost.ID); err != nil {
		return err
	}

	if updatePost.ID == 0 {
		return errors.New("Không tìm thấy bài viết cần cập nhật với ID: " + string(domainPost.ID))
	}

	if domainPost.Title != "" {
		updatePost.Title = domainPost.Title
	}

	if domainPost.Title != "" {
		updatePost.Title = domainPost.Title
		updatePost.Slug = uc.service.GenerateSlug(domainPost.Title)
	}

	if domainPost.Description != "" {
		updatePost.Description = domainPost.Description
	}

	if isRepost {
		const repostCooldown = 7 * 24 * time.Hour

		if time.Since(updatePost.CreatedAt) >= repostCooldown {
			// Được phép repost
			updatePost.CreatedAt = time.Now()
		} else {
			remaining := repostCooldown - time.Since(updatePost.CreatedAt)
			// Chưa đủ thời gian để repost
			return errors.New("Bạn chỉ có thể repost sau " + string(remaining.Truncate(time.Minute)) + " nữa")
		}
	}

	if domainPost.Images != nil {
		var images []string

		//resize ảnh
		for _, image := range domainPost.Images {
			formatedImage, err := helpers.ProcessImageBase64(image, uint(enums.PostImageWidth), uint(enums.PostImageHeight), 75, helpers.FormatJPEG)
			if err != nil {
				return errors.New("Không thể format ảnh:" + err.Error())
			}

			images = append(images, formatedImage)
		}

		updatePost.Images = images
	}

	updatePost.Status = domainPost.Status

	if domainPost.Status == int8(enums.PostStatusApproved) {
		updatePost.CreatedAt = time.Now()
	}

	if err := uc.repo.Update(ctx, &updatePost); err != nil {
		return err
	}

	return nil
}
