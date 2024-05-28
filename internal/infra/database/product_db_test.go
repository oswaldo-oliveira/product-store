package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/oswaldo-oliveira/productstore/internal/entities"
	"github.com/oswaldo-oliveira/productstore/test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewCreateProduct(t *testing.T) {
	product, _ := entities.NewProduct("Product 1", 10)
	productDB := repositories.NewProduct([]entities.Product{})

	productDB.Create(product)

	productFound := productDB.DB[0]

	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	productDB := repositories.NewProduct([]entities.Product{})

	for i := 1; i < 25; i++ {
		product, _ := entities.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		productDB.Create(product)
	}

	products, _ := productDB.FindAll(0, 0, "desc")
	assert.Len(t, products, 24)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 24", products[23].Name)

	products, _ = productDB.FindAll(1, 10, "asc")
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)
}

func TestFindProductById(t *testing.T) {
	product, _ := entities.NewProduct("Product 1", 10)
	productDB := repositories.NewProduct([]entities.Product{})

	productDB.Create(product)
	productFound, _ := productDB.FindByID(product.ID.String())

	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Product 1", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	product, _ := entities.NewProduct("Product 1", 10)
	productDB := repositories.NewProduct([]entities.Product{})

	productDB.Create(product)

	productDB.Delete(product.ID.String())
	productFound, _ := productDB.FindByID(product.ID.String())

	assert.Nil(t, productFound)
}

func TestUpdateProduct(t *testing.T) {
	product, _ := entities.NewProduct("Product 1", 10)
	productDB := repositories.NewProduct([]entities.Product{})

	productDB.Create(product)

	product.Name = "Product 2"
	productDB.Update(product)
	productFound, _ := productDB.FindByID(product.ID.String())

	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Product 2", productFound.Name)
}
