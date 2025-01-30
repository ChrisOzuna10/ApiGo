package application

import (
	"api/src/products/domain"
)

type UpdateProduct struct {
	db domain.IProduct
}

func NewUpdateProduct(db domain.IProduct) *UpdateProduct {
	return &UpdateProduct{db: db}
}

func (up *UpdateProduct) Update(id int32, name string, price float32) error {
	return up.db.Update(id, name, price)
}
