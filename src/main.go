package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"taco-api/src/infra/gorm"
	gormRepositories "taco-api/src/infra/gorm/repositories"
	config "taco-api/src/infra/viper"
	"taco-api/src/interface/controllers"
)

func main() {
	configs := config.Init()

	gormConnection := gorm.NewConnection(configs.DSN)

	categoryRepo := gormRepositories.NewGormCategoryRepository(gormConnection.Connection)
	createCategoryController := controllers.NewCreateCategoryController(categoryRepo)

	categoryCreated, error := createCategoryController.Execute("Teste do 6")

	if error != nil {
		log.Fatal(error.Error())
	}

	fmt.Println(categoryCreated)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Print("Server running at 3001!")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
