package repositories

import (
	"log"
	"taco-api/src/domain"
	"taco-api/src/infra/gorm/models"

	"gorm.io/gorm"
)

type GormFoodRepository struct {
	connection *gorm.DB
}

func NewGormFoodRepository(connection *gorm.DB) GormFoodRepository {
	return GormFoodRepository{connection: connection}
}

func (r GormFoodRepository) Save(food *domain.Food) *domain.Food {
	gormFood := models.GormFoodModel{
		Uuid:         food.ID,
		Name:         food.Name,
		CategoryUuid: food.Category.ID,
		Kcal:         food.Kcal,
		Protein:      food.Protein,
		Carbohydrate: food.Carbohydrate,
		Lipid:        food.Lipid,
		Fiber:        food.Fiber,
		Ashes:        food.Ashes,
		Calcium:      food.Calcium,
		Magnesium:    food.Magnesium,
		Manganese:    food.Manganese,
		Phosphorus:   food.Phosphorus,
		Iron:         food.Iron,
		Sodium:       food.Sodium,
		Potassium:    food.Potassium,
		Zinc:         food.Zinc,
	}

	result := r.connection.Create(&gormFood)

	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}

	return food
}
