package controllers

import (
	"taco-api/src/application/repositories"
	"taco-api/src/application/usecases"
	"taco-api/src/domain"
	"taco-api/src/interface/dtos"
)

type CreateFoodController struct {
	usecase usecases.CreateFoodUsecase
}

func NewCreateFoodController(
	categoryRepo repositories.CategoryRepository,
	foodRepo repositories.FoodRepository,
) CreateFoodController {
	return CreateFoodController{
		usecase: usecases.NewCreateFoodUsecase(categoryRepo, foodRepo),
	}
}

func (c CreateFoodController) Execute(name string, categoryId string, nutritionalInformations domain.NutritionalInformations) (dtos.FoodDTO, error) {
	food, error := c.usecase.Perform(name, categoryId, nutritionalInformations)

	if error != nil {
		return dtos.FoodDTO{}, error
	}

	return dtos.NewFoodDTO(food), nil
}
