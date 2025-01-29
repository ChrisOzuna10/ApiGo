package application

import (
	"api/src/products/domain"
    "api/src/products/domain/entities"
)

type ViewProduct struct {
	db domain.IProduct
}

func NewViewProduct(db domain.IProduct) *ViewProduct {
    return &ViewProduct{db: db}
}

func (vp *ViewProduct) Execute() ([]entities.Product, error) {
    return vp.db.GetAll()
}