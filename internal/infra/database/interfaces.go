package database

import "github.com/oswaldo-oliveira/productstore/internal/entities"

type UserInterface interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}
