package postdto

import (
	"final_project/internal/domain/post"
)

// DTO → Domain
func CreateDTOToDomain(createPostDTO CreatePostRequest) post.CreatePost {
	return post.CreatePost{
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

// DTO → Domain
func UpdateDTOToDomain(updatePostDTO UpdatePostRequest) post.Post {
	return post.Post{
		Title:  updatePostDTO.Title,
		Info:   updatePostDTO.Info,
		Status: int8(updatePostDTO.Status),
		Images: updatePostDTO.Images,
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
