package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/oswaldo-oliveira/productstore/configs"
	"github.com/oswaldo-oliveira/productstore/internal/entities"
	"github.com/oswaldo-oliveira/productstore/internal/infra/database"
	"github.com/oswaldo-oliveira/productstore/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Product{}, &entities.User{})
	productDB := database.NewProduct(db)
	userDB := database.NewUser(db)
	productHandler := handlers.NewProductHandler(productDB)
	usersHandler := handlers.NewUserHandler(userDB, conf.TokenAuth, conf.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(conf.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", usersHandler.CreateUser)
		r.Post("/login", usersHandler.Login)
	})

	http.ListenAndServe(":8000", r)
}
