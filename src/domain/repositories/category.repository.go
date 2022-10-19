package repositories

import "taco-api/src/domain/entities"

type CategoryRepository interface {
	Save(category entities.Category) (entities.Category, error)
	GetById(id string) entities.Category
	GetAll() []entities.Category
	Update(category entities.Category) entities.Category
	DeleteById(id string)
}
