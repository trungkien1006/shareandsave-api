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

func (uc *UseCase) CreatePost(ctx context.Context, post *post.CreatePost) error {
	var (
		author user.User
		tags   []string
	)

	itemIDs := make(map[uint]uint, 0)

	if err := uc.userRepo.GetCommonUserByID(ctx, &author, int(post.AuthorID)); err != nil {
		return err
	}

	if post.Info != "" {
		postContent, err := uc.service.GenerateContent(post.Info)
		if err != nil {
			return errors.New("Lỗi khi tạo content từ info:" + err.Error())
		}

		post.Content = postContent
	} else {
		post.Content = "{}"
		post.Info = "{}"
	}

	post.AuthorID = author.ID
	post.Status = int8(enums.PostStatusPending) // Mặc định trạng thái là Pending
	post.Slug = uc.service.GenerateSlug(post.Title)
	post.AuthorName = author.FullName

	//resize ảnh
	for index, image := range post.Images {
		formatedImage, err := helpers.ProcessImageBase64(image, uint(enums.PostImageWidth), uint(enums.PostImageHeight), 75, helpers.FormatJPEG)
		if err != nil {
			return errors.New("Không thể format ảnh:" + err.Error())
		}

		post.Images[index] = formatedImage
	}

	for key, oldItem := range post.OldItems {
		var item item.Item

		err := uc.itemRepo.GetByID(ctx, &item, oldItem.ItemID)
		if err != nil {
			return err
		}

		if oldItem.Image == "" {
			post.OldItems[key].Image = item.Image
		} else {
			strBase64Image, err := helpers.ProcessImageBase64(oldItem.Image, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)

			if err != nil {
				return err
			}

			post.OldItems[key].Image = strBase64Image
		}

		itemIDs[oldItem.ItemID] = oldItem.ItemID
	}

	for key, newItem := range post.NewItems {
		item := item.Item{
			CategoryID: newItem.CategoryID,
			Name:       newItem.Name,
		}

		if newItem.Image == "" {
			base64, err := helpers.ImageToBase64(os.Getenv("IMAGE_PATH") + "/item.png")
			if err != nil {
				return err
			}

			strBase64Image, err := helpers.ProcessImageBase64(base64, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)
			if err != nil {
				return err
			}

			item.Image = strBase64Image
		} else {
			strBase64Image, err := helpers.ProcessImageBase64(newItem.Image, uint(enums.ItemImageWidth), uint(enums.ItemImageHeight), 75, helpers.FormatJPEG)

			if err != nil {
				return err
			}

			item.Image = strBase64Image
		}

		if err := uc.itemRepo.Save(ctx, &item); err != nil {
			return err
		}

		post.NewItems[key].ItemID = item.ID

		itemIDs[item.ID] = item.ID
	}

	if err := uc.categoryRepo.GetCategoryNameByItemIDs(ctx, itemIDs, &tags); err != nil {
		return err
	}

	post.Tag = tags

	//create post
	if err := uc.repo.Save(ctx, post); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdatePost(ctx context.Context, domainPost *post.Post) error {
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

	if err := uc.repo.Update(ctx, &updatePost); err != nil {
		return err
	}

	return nil
}
