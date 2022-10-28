package models

import "gorm.io/gorm"

type GormCategoryModel struct {
	gorm.Model
	Uuid  string          `gorm:"unique"`
	Name  string          `gorm:"unique"`
	Foods []GormFoodModel `gorm:"foreignKey:CategoryUuid;references:Uuid"`
}

func (GormCategoryModel) TableName() string {
	return "categories"
}
