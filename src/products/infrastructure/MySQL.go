package infrastructure

import (
	"api/src/core"
	"api/src/products/domain"
	"api/src/products/domain/entities"
)

type MySQLRepositoryProducts struct {}

func NewMySQLRepositoryProducts() domain.IProduct {
    return &MySQLRepositoryProducts{}
}

func (repo *MySQLRepositoryProducts) Save(name string, price float32) error {
	product := entities.Product{Name: name, Price: price}
	result:= core.BD.Create(&product)
	return result.Error
}

func (repo *MySQLRepositoryProducts) GetAll() ([]entities.Product, error) {
	var products []entities.Product
	result := core.BD.Find(&products)
	return products, result.Error
}

func (repo *MySQLRepositoryProducts) Delete(id int32) error {
	result := core.BD.Where ("id_product = ?", id).Delete(&entities.Product{})
    return result.Error
}