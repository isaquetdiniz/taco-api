package usecases

import (
	"log"
	"taco-api/src/application/exceptions"
	"taco-api/src/application/repositories"
	"taco-api/src/domain"
)

type CreateCategoryUsecase struct {
	repository repositories.CategoryRepository
}

func NewCreateCategoryUsecase(
	categoryRepo repositories.CategoryRepository,
) CreateCategoryUsecase {
	return CreateCategoryUsecase{repository: categoryRepo}
}

func (c CreateCategoryUsecase) Perform(name string) (domain.Category, error) {
	_, exists, error := c.repository.GetByName(name)

	if error != nil {
		log.Fatal(error.Error())
	}

	if exists {
		return domain.Category{}, exceptions.CategoryAlreadyExists
	}

	category := domain.NewCategory(name)

	categoryCreated := c.repository.Save(category)

	return *categoryCreated, nil
}
