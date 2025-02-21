package useCaseProduct

import (
	"api/src/products/domain"
)

type DeleteProduct struct {
	db domain.IProduct
}

func NewDeleteProduct(db domain.IProduct) *DeleteProduct {
	return &DeleteProduct{db: db}
}

func (dp *DeleteProduct) Execute(id int32) error {
	return dp.db.Delete(id)
}
