package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oswaldo-oliveira/productstore/configs"
	"github.com/oswaldo-oliveira/productstore/internal/entities"
	"github.com/oswaldo-oliveira/productstore/internal/infra/database"
	"github.com/oswaldo-oliveira/productstore/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entities.Product{}, &entities.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)

	http.ListenAndServe(":8000", r)
}
