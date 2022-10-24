package models

import "gorm.io/gorm"

type GormCategoryModel struct {
	gorm.Model
	Uuid string `gorm:"unique"`
	Name string `gorm:"unique"`
}

func (GormCategoryModel) TableName() string {
	return "categories"
}
