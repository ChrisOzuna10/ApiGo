package application

import "api/src/products/domain"

type CreateProduct struct {
	db domain.IProduct
}

func NewCreateProduct(db domain.IProduct) *CreateProduct {
    return &CreateProduct{db: db}
}

func (cp *CreateProduct) Execute(name string, price float32) error {
	return cp.db.Save(name, price)
}