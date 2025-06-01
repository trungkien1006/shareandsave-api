package categorydto

import "final_project/internal/domain/category"

// Domain -> DTO
func CateDomainToDTO(domainCate category.Category) CategoryDTO {
	return CategoryDTO{
		ID:   domainCate.ID,
		Name: domainCate.Name,
	}
}
