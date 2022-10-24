package repositories

import "taco-api/src/domain"

type FoodRepository interface {
	Repository[domain.Food]
}
