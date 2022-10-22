package gorm

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormConnection struct {
	connection *gorm.DB
}

func NewGormConnection(dsn string) *GormConnection {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return &GormConnection{connection: db}
}

func (g GormConnection) Close() {
	db, err := g.connection.DB()

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Close()
}
