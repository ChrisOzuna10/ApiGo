package domain

import "api/src/domain/entities"

type IProduct interface {
	Save(name string, price float32) error
	GetAll() ([]entities.Product, error)
}