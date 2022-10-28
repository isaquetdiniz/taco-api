package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"taco-api/src/domain"
	"taco-api/src/infra/gorm"
	gormRepositories "taco-api/src/infra/gorm/repositories"
	config "taco-api/src/infra/viper"
	"taco-api/src/interface/controllers"
)

func main() {
	configs := config.Init()

	gormConnection := gorm.NewConnection(configs.DSN)

	categoryRepo := gormRepositories.NewGormCategoryRepository(gormConnection.Connection)
	foodRepo := gormRepositories.NewGormFoodRepository(gormConnection.Connection)

	createCategoryController := controllers.NewCreateCategoryController(categoryRepo)
	createFoodController := controllers.NewCreateFoodController(categoryRepo, foodRepo)

	categoryCreated, error1 := createCategoryController.Execute("Teste do 6")

	if error1 != nil {
		log.Fatal(error1.Error())
	}

	fmt.Println(categoryCreated)

	nt := domain.NutritionalInformations{
		Protein:      10,
		Kcal:         1000,
		Carbohydrate: 1000,
		Lipid:        1000,
	}

	foodCreated, error2 := createFoodController.Execute("Teste de comida", categoryCreated.ID, nt)

	if error2 != nil {
		log.Fatal(error2.Error())
	}

	fmt.Println(foodCreated)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Print("Server running at 3001!")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
