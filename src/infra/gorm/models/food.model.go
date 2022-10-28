package models

import "gorm.io/gorm"

type GormFoodModel struct {
	gorm.Model
	Uuid         string `gorm:"unique"`
	Name         string `gorm:"unique"`
	CategoryUuid string
	Kcal         int
	Protein      int
	Carbohydrate int
	Lipid        int
	Fiber        int
	Ashes        int
	Calcium      int
	Magnesium    int
	Manganese    int
	Phosphorus   int
	Iron         int
	Sodium       int
	Potassium    int
	Zinc         int
}

func (GormFoodModel) TableName() string {
	return "foods"
}
