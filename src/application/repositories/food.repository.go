package repositories

import "taco-api/src/domain"

type FoodRepository interface {
	Save(*domain.Food) *domain.Food
}
