package repositories

import "taco-api/src/domain/entities"

type CategoryRepository interface {
	Save(*entities.Category) *entities.Category
	GetByName(name string) (entities.Category, bool, error)
}
