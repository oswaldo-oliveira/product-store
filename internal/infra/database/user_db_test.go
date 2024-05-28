package database

import (
	"testing"

	"github.com/oswaldo-oliveira/productstore/internal/entities"
	"github.com/oswaldo-oliveira/productstore/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user, _ := entities.NewUser("John Doe", "john.doe@email.com", "123456")
	userDB := repositories.NewUser([]entities.User{})

	userDB.Create(user)

	userFound := userDB.DB[0]

	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)

}

func TestFindByEmail(t *testing.T) {
	user, _ := entities.NewUser("John Doe", "john.doe@email.com", "123456")
	userDB := repositories.NewUser([]entities.User{})

	userDB.Create(user)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
