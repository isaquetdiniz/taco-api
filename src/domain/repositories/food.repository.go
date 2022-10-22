package repositories

import "taco-api/src/domain/entities"

type FoodRepository interface {
	Repository[entities.Food]
}
