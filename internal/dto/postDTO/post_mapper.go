package postdto

import (
	"final_project/internal/domain/post"
	"final_project/internal/pkg/enums"
)

// Domain -> DTO
func PostWithCountDomainToDTO(domain post.PostWithCount) PostWithCountDTO {
	DTOTag := make([]string, 0)
	DTOImage := make([]string, 0)

	for _, value := range domain.Tag {
		DTOTag = append(DTOTag, value)
	}

	for _, value := range domain.Images {
		DTOImage = append(DTOImage, value)
	}

	return PostWithCountDTO{
		ID:               domain.ID,
		AuthorID:         domain.AuthorID,
		AuthorName:       domain.AuthorName,
		AuthorAvatar:     domain.AuthorAvatar,
		Type:             domain.Type,
		Slug:             domain.Slug,
		Title:            domain.Title,
		Description:      domain.Description,
		Content:          domain.Content,
		Info:             domain.Info,
		Status:           domain.Status,
		Images:           DTOImage,
		CreatedAt:        domain.CreatedAt,
		Tag:              DTOTag,
		InterestCount:    domain.InterestCount,
		ItemCount:        domain.ItemCount,
		CurrentItemCount: domain.CurrentItemCount,
	}
}

// Domain -> DTO
func DetailPostDomainToDTO(domain post.DetailPost) DetailPostDTO {
	DTOInterest := make([]InterestDTO, 0)
	DTOPostItem := make([]DetailPostItemDTO, 0)
	DTOTag := make([]string, 0)
	DTOImage := make([]string, 0)

	for _, value := range domain.Interest {
		DTOInterest = append(DTOInterest, InterestDTO{
			ID:         value.ID,
			UserID:     value.UserID,
			UserName:   value.UserName,
			UserAvatar: value.UserAvatar,
			PostID:     value.PostID,
			Status:     value.Status,
		})
	}

	for _, value := range domain.Items {
		DTOPostItem = append(DTOPostItem, DetailPostItemDTO{
			ID:              value.ID,
			ItemID:          value.ItemID,
			CategoryID:      value.CategoryID,
			CategoryName:    value.CategoryName,
			Image:           value.Image,
			Name:            value.Name,
			Quantity:        value.Quantity,
			CurrentQuantity: value.CurrentQuantity,
		})
	}

	for _, value := range domain.Tag {
		DTOTag = append(DTOTag, value)
	}

	for _, value := range domain.Images {
		DTOImage = append(DTOImage, value)
	}

	return DetailPostDTO{
		ID:          domain.ID,
		AuthorID:    domain.AuthorID,
		AuthorName:  domain.AuthorName,
		Type:        domain.Type,
		Slug:        domain.Slug,
		Title:       domain.Title,
		Description: domain.Description,
		Content:     domain.Content,
		Info:        domain.Info,
		Status:      domain.Status,
		Images:      DTOImage,
		CreatedAt:   domain.CreatedAt,
		Tag:         DTOTag,
		Interest:    DTOInterest,
		Items:       DTOPostItem,
	}
}

// DTO → Domain
func CreateDTOToDomain(createPostDTO CreatePostRequest) post.CreatePost {
	var domainOldItems []post.OldItemsInPost

	for _, value := range createPostDTO.OldItems {
		domainOldItems = append(domainOldItems, OldItemsDTOToDomain(value))
	}

	var domainNewItems []post.NewItemsInPost

	for _, value := range createPostDTO.NewItems {
		domainNewItems = append(domainNewItems, NewItemsDTOToDomain(value))
	}

	return post.CreatePost{
		// FullName:    createPostDTO.FullName,
		// Email:       createPostDTO.Email,
		// PhoneNumber: createPostDTO.PhoneNumber,
		Type:        int(createPostDTO.Type),
		Title:       createPostDTO.Title,
		Info:        createPostDTO.Info,
		Description: createPostDTO.Description,
		Images:      createPostDTO.Images,
		OldItems:    domainOldItems,
		NewItems:    domainNewItems,
	}
}

// Domain -> DTO
func CreateDomainToDTO(domain post.CreatePost) CreatePostRequest {
	var domainOldItems []OldItemsPost

	for _, value := range domain.OldItems {
		domainOldItems = append(domainOldItems, OldItemsDomainToDTO(value))
	}

	var domainNewItems []NewItemsPost

	for _, value := range domain.NewItems {
		domainNewItems = append(domainNewItems, NewItemsDomainToDTO(value))
	}

	return CreatePostRequest{
		// FullName:    domain.FullName,
		// Email:       domain.Email,
		// PhoneNumber: domain.PhoneNumber,
		Type:     enums.PostType(domain.Type),
		Title:    domain.Title,
		Info:     domain.Info,
		Images:   domain.Images,
		OldItems: domainOldItems,
		NewItems: domainNewItems,
	}
}

// OldItems DTO -> Domain
func OldItemsDTOToDomain(item OldItemsPost) post.OldItemsInPost {
	return post.OldItemsInPost{
		ItemID:   item.ItemID,
		Image:    item.Image,
		Quantity: item.Quantity,
	}
}

// Domain -> OldItems DTO
func OldItemsDomainToDTO(item post.OldItemsInPost) OldItemsPost {
	return OldItemsPost{
		ItemID:   item.ItemID,
		Image:    item.Image,
		Quantity: item.Quantity,
	}
}

// NewItems DTO -> Domain
func NewItemsDTOToDomain(item NewItemsPost) post.NewItemsInPost {
	return post.NewItemsInPost{
		CategoryID: item.CategoryID,
		Image:      item.Image,
		Name:       item.Name,
		Quantity:   item.Quantity,
	}
}

// Domain -> NewItems DTO
func NewItemsDomainToDTO(item post.NewItemsInPost) NewItemsPost {
	return NewItemsPost{
		CategoryID: item.CategoryID,
		Image:      item.Image,
		Name:       item.Name,
		Quantity:   item.Quantity,
	}
}

// DTO → Domain
func UpdateDTOToDomain(updatePostDTO UpdatePostRequest) post.Post {
	return post.Post{
		Title:       updatePostDTO.Title,
		Description: updatePostDTO.Description,
		Status:      int8(updatePostDTO.Status),
		Images:      updatePostDTO.Images,
	}
}

// Domain → DTO
func DomainAdminPostToDTO(domainPost post.Post) AdminPostDTO {
	return AdminPostDTO{
		ID:         domainPost.ID,
		AuthorName: domainPost.AuthorName,
		Type:       domainPost.Type,
		Title:      domainPost.Title,
		Status:     domainPost.Status,
		CreatedAt:  domainPost.CreatedAt,
		IsInterest: domainPost.IsInterested,
	}
}

func CreatePostDomainToDTO(domainPost post.CreatePost) PostDTO {
	return PostDTO{
		ID:         domainPost.ID,
		AuthorName: domainPost.AuthorName,
		Content:    domainPost.Content,
		Slug:       domainPost.Slug,
		Type:       domainPost.Type,
		Title:      domainPost.Title,
		Status:     domainPost.Status,
		Images:     domainPost.Images,
	}
}
