package repositories

import "taco-api/src/domain/entities"

type FoodRepository interface {
	Save(food entities.Food) entities.Food
	GetById(id string) entities.Food
	Update(food entities.Food) entities.Food
	DeleteById(id string)
}
