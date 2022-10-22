package repositories

import (
	"log"
	"taco-api/src/domain/entities"
	"taco-api/src/infra/gorm/models"

	"gorm.io/gorm"
)

type GormCategoryRepository struct {
	connection *gorm.DB
}

func NewGormCategoryRepository(connection *gorm.DB) GormCategoryRepository {
	return GormCategoryRepository{connection: connection}
}

func (r GormCategoryRepository) Save(category *entities.Category) *entities.Category {

	gormCategory := models.GormGategoryModel{
		Uuid: category.ID,
		Name: category.Name,
	}

	result := r.connection.Create(&gormCategory)

	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}

	return category
}
