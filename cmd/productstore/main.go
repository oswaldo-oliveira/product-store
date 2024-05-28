package main

import (
	"net/http"

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

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
