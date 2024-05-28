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

func (p *Product) FindAll(page, limit int, sort string) ([]entities.Product, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		return p.DB[page*limit : page*limit+limit], nil
	}
	return p.DB, nil
}

func (p *Product) Update(product *entities.Product) error {
	for index, prod := range p.DB {
		if prod.ID == product.ID {
			p.DB[index] = *product
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
