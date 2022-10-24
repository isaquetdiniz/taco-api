package dtos

import "taco-api/src/domain/entities"

type CategoryDTO = EntityDTO

func NewCategoryDTO(category entities.Category) CategoryDTO {
	return CategoryDTO{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
