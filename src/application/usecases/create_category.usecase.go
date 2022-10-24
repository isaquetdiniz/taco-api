package usecases

import (
	"log"
	"taco-api/src/application/exceptions"
	"taco-api/src/domain/entities"
	"taco-api/src/domain/repositories"
)

type CreateCategoryUsecase struct {
	repository repositories.CategoryRepository
}

func NewCreateCategoryUsecase(
	categoryRepo repositories.CategoryRepository,
) CreateCategoryUsecase {
	return CreateCategoryUsecase{repository: categoryRepo}
}

func (c CreateCategoryUsecase) Perform(name string) (entities.Category, error) {
	_, exists, error := c.repository.GetByName(name)

	if error != nil {
		log.Print(error.Error())
	}

	if exists {
		return entities.Category{}, exceptions.CategoryAlreadyExists
	}

	category := entities.NewCategory(name)

	categoryCreated := c.repository.Save(category)

	return *categoryCreated, nil
}
