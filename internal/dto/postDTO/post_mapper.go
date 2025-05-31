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
func DomainToDTO(domainPost post.Post) AdminPostDTO {
	return AdminPostDTO{
		ID:         domainPost.ID,
		AuthorName: domainPost.AuthorName,
		Type:       domainPost.Type,
		Slug:       domainPost.Slug,
		Title:      domainPost.Title,
		Content:    domainPost.Content,
		Status:     domainPost.Status,
	}
}
