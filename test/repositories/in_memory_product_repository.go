package repositories

import (
	"errors"

	"github.com/oswaldo-oliveira/productstore/internal/entities"
)

type Product struct {
	DB []entities.Product
}

func NewProduct(db []entities.Product) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entities.Product) error {
	p.DB = append(p.DB, *product)
	return nil
}

func (p *Product) FindByID(id string) (*entities.Product, error) {
	for _, product := range p.DB {
		if product.ID.String() == id {
			return &product, nil
		}
	}
	return nil, errors.New("Product not found")
}

func (p *Product) Update(product *entities.Product) error {
	for _, prod := range p.DB {
		if prod.ID == product.ID {
			prod = *product
			return nil
		}
	}

	return errors.New("Product not found")
}

func (p *Product) Delete(id string) error {
	for index, prod := range p.DB {
		if prod.ID.String() == id {
			p.DB = append(p.DB[:index], p.DB[index+1:]...)
			return nil
		}
	}

	return errors.New("Product not found")
}
