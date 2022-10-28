package dtos

import "taco-api/src/domain"

type FoodDTO struct {
	EntityDTO
	CategoryDTO
	domain.NutritionalInformations
}

func NewFoodDTO(food domain.Food) FoodDTO {
	entityDTO := EntityDTO{
		ID:        food.ID,
		Name:      food.Name,
		CreatedAt: food.CreatedAt,
		UpdatedAt: food.UpdatedAt,
	}

	categoryDTO := NewCategoryDTO(food.Category)

	return FoodDTO{
		EntityDTO:               entityDTO,
		CategoryDTO:             categoryDTO,
		NutritionalInformations: *food.NutritionalInformations,
	}
}
