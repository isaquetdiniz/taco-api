package gorm

import (
	"log"
	"taco-api/src/infra/gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormConnection struct {
	Connection *gorm.DB
}

func NewConnection(dsn string) *GormConnection {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.GormCategoryModel{})
	db.AutoMigrate(&models.GormFoodModel{})

	return &GormConnection{Connection: db}
}

func (g GormConnection) Close() {
	db, err := g.Connection.DB()

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Close()
}
