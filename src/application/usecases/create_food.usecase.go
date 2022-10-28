package usecases

import (
	"log"
	"taco-api/src/application/exceptions"
	"taco-api/src/application/repositories"
	"taco-api/src/domain"
)

type CreateFoodUsecase struct {
	categoryRepository repositories.CategoryRepository
	foodRepository     repositories.FoodRepository
}

func NewCreateFoodUsecase(
	categoryRepo repositories.CategoryRepository,
	foodRepo repositories.FoodRepository,
) CreateFoodUsecase {

	return CreateFoodUsecase{
		categoryRepository: categoryRepo,
		foodRepository:     foodRepo,
	}
}

func (f CreateFoodUsecase) Perform(name string, categoryId string, nutritionalInformations domain.NutritionalInformations) (domain.Food, error) {
	categoryFound, existsCategory, error := f.categoryRepository.GetById(categoryId)

	if error != nil {
		log.Fatal(error.Error())
	}

	if !existsCategory {
		return domain.Food{}, exceptions.CategoryNotFound
	}

	food := domain.NewFood(name, categoryFound, &nutritionalInformations)

	foodCreated := f.foodRepository.Save(food)

	return *foodCreated, nil
}
