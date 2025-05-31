package postdto

import (
	"final_project/internal/domain/post"
)

// DTO → Domain
func CreateDTOToDomain(createPostDTO CreatePostRequest) post.Post {
	return post.Post{
		AuthorID:    createPostDTO.AuthorID,
		FullName:    createPostDTO.FullName,
		Email:       createPostDTO.Email,
		PhoneNumber: createPostDTO.PhoneNumber,
		Type:        int(createPostDTO.Type),
		Title:       createPostDTO.Title,
		Info:        createPostDTO.Info,
		Images:      createPostDTO.Images,
	}
}

// Domain → DTO
func DomainAdminPostToDTO(domainPost post.AdminPost) AdminPostDTO {
	return AdminPostDTO{
		ID:         domainPost.ID,
		AuthorName: domainPost.AuthorName,
		Type:       domainPost.Type,
		Title:      domainPost.Title,
		Status:     domainPost.Status,
		CreatedAt:  domainPost.CreateAt,
	}
}

func DomainToDTO(domainPost post.Post) PostDTO {
	return PostDTO{
		ID:         domainPost.ID,
		AuthorName: domainPost.AuthorName,
		Content:    domainPost.Content,
		Slug:       domainPost.Slug,
		Type:       domainPost.Type,
		Title:      domainPost.Title,
		Status:     domainPost.Status,
	}
}
