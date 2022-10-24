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

	gormCategory := models.GormCategoryModel{
		Uuid: category.ID,
		Name: category.Name,
	}

	result := r.connection.Create(&gormCategory)

	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}

	return category
}

func (r GormCategoryRepository) GetByName(name string) (entities.Category, bool, error) {
	var gormCategory models.GormCategoryModel

	result := r.connection.First(&gormCategory, "name = ?", name)

	if result.RowsAffected == 0 {
		return entities.Category{}, false, nil
	}

	category := entities.NewCategory(
		name,
		entities.WithID(gormCategory.Uuid),
		entities.WithCreatedAt(gormCategory.CreatedAt),
		entities.WithUpdatedAt(gormCategory.UpdatedAt),
	)

	return *category,
		true,
		nil
}
