package repositories

import "taco-api/src/domain"

type CategoryRepository interface {
	Save(*domain.Category) *domain.Category
	GetByName(name string) (domain.Category, bool, error)
}
