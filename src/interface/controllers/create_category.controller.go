package controllers

import (
	"log"
	"taco-api/src/application/usecases"
	"taco-api/src/domain/repositories"
	"taco-api/src/interface/dtos"
)

type CreateCategoryController struct {
	usecase usecases.CreateCategoryUsecase
}

func NewCreateCategoryUsecase(
	categoryRepo repositories.CategoryRepository
  ) CreateCategoryController {
	return CreateCategoryController{
		usecase: usecases.NewCreateCategoryUsecase(),
	}
}

func (c CreateCategoryController) Execute(name string) dtos.EntityDTO {
  category, error := c.usecase.Perform(name)

  if error != nil {
    log.Fatal(error.Error())
  }

  return dtos.EntityDTO{
    ID: category.ID,
    Name: category.Name,
    CreatedAt: category.CreatedAt,
    UpdatedAt: category.UpdatedAt,
  }
}
