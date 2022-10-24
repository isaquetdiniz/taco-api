package dtos

import "taco-api/src/domain"

type CategoryDTO = EntityDTO

func NewCategoryDTO(category domain.Category) CategoryDTO {
	return CategoryDTO{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
