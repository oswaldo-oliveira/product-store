package database

import (
	"github.com/oswaldo-oliveira/productstore/internal/entities"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entities.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByID(id string) (*entities.Product, error) {
	var product entities.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Update(product *entities.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	_, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(id).Error
}
