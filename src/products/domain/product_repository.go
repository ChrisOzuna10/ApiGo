package domain

import "api/src/products/domain/entities"

type IProduct interface {
	Save(name string, price float32) error
	GetAll() ([]entities.Product, error)
	GetById(id int32) (*entities.Product, error)
	Delete(id int32) error
	Update(id int32, name string, price float32) error
}
