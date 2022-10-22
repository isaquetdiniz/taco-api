package models

import "gorm.io/gorm"

type GormGategoryModel struct {
	gorm.Model
	Uuid string
	Name string
}
