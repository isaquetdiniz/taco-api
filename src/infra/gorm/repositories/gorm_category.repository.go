package repositories

import (
	"fmt"
	"log"
	"taco-api/src/domain"
	"taco-api/src/infra/gorm/models"

	"gorm.io/gorm"
)

type GormCategoryRepository struct {
	connection *gorm.DB
}

func NewGormCategoryRepository(connection *gorm.DB) GormCategoryRepository {
	return GormCategoryRepository{connection: connection}
}

func (r GormCategoryRepository) Save(category *domain.Category) *domain.Category {

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

func (r GormCategoryRepository) GetByName(name string) (domain.Category, bool, error) {
	var gormCategory models.GormCategoryModel

	result := r.connection.First(&gormCategory, "name = ?", name)

	if result.RowsAffected == 0 {
		return domain.Category{}, false, nil
	}

	category := domain.NewCategory(
		name,
		domain.WithID(gormCategory.Uuid),
		domain.WithCreatedAt(gormCategory.CreatedAt),
		domain.WithUpdatedAt(gormCategory.UpdatedAt),
	)

	return *category,
		true,
		nil
}

func (r GormCategoryRepository) GetById(id string) (domain.Category, bool, error) {
	var gormCategory models.GormCategoryModel

	result := r.connection.First(&gormCategory, "uuid = ?", id)

	if result.RowsAffected == 0 {
		return domain.Category{}, false, nil
	}

	fmt.Println(result)

	category := domain.NewCategory(
		result.Name(),
		domain.WithID(gormCategory.Uuid),
		domain.WithCreatedAt(gormCategory.CreatedAt),
		domain.WithUpdatedAt(gormCategory.UpdatedAt),
	)

	return *category,
		true,
		nil
}
