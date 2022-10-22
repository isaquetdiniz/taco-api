package repositories

import "taco-api/src/domain/entities"

type CategoryRepository interface {
	Repository[entities.Category]
	GetByName(name string) (*entities.Category, bool, error)
}
