package useCaseProduct

import (
	"api/src/products/domain"
	"api/src/products/domain/entities"
)

type ViewByIdProduct struct {
	db domain.IProduct
}

func NewViewByIdProduct(db domain.IProduct) *ViewByIdProduct {
	return &ViewByIdProduct{db: db}
}

func (vp *ViewByIdProduct) GetById(id int32) (*entities.Product, error) {
	return vp.db.GetById(id)
}
