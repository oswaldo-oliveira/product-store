package repositories

import (
	"github.com/oswaldo-oliveira/productstore/internal/entities"
)

type User struct {
	DB []entities.User
}

func NewUser(db []entities.User) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entities.User) error {
	u.DB = append(u.DB, *user)
	return nil
}

func (u *User) FindByEmail(email string) (*entities.User, error) {
	for _, user := range u.DB {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}
