package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"taco-api/src/domain/entities"
	config "taco-api/src/infra/viper"
)

func main() {
	configs := config.Init()

	fmt.Println(configs)

	//teste := gorm.NewGormConnection("")

	//fmt.Println(teste)

	category := entities.NewCategory("Cereal")

	nutritionalInformations := &entities.NutritionalInformations{
		Protein:      10,
		Carbohydrate: 20,
		Lipid:        5,
	}

	f := entities.NewFood(
		"Arroz",
		*category,
		nutritionalInformations,
	)

	fmt.Println(f)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Print("Server running at 3001!")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
