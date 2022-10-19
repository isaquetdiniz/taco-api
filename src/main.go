package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"taco-api/src/domain/entities"
	"time"
)

func main() {
	c := entities.Category{
		Entity: entities.Entity{Id: "id", Name: "Cereal", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local()}}

	fmt.Println(c)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Print("Server running at 3001!")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
