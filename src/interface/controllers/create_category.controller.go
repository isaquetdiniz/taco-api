package controllers

import (
	"taco-api/src/application/usecases"
	"taco-api/src/domain/repositories"
	"taco-api/src/interface/dtos"
)

type CreateCategoryController struct {
	usecase usecases.CreateCategoryUsecase
}

func NewCreateCategoryController(
	categoryRepo repositories.CategoryRepository,
) CreateCategoryController {
	return CreateCategoryController{
		usecase: usecases.NewCreateCategoryUsecase(categoryRepo),
	}
}

func (c CreateCategoryController) Execute(name string) (dtos.EntityDTO, error) {
	category, error := c.usecase.Perform(name)

	if error != nil {
		return dtos.CategoryDTO{}, error
	}

	return dtos.NewCategoryDTO(category), nil
}
