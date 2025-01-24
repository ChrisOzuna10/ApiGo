package infrastructure

import (
	"api/src/core"
	"api/src/domain"
	"api/src/domain/entities"
)

type MySQLRepository struct {}

func NewMySQLRepository() domain.IProduct {
    return &MySQLRepository{}
}

func (repo *MySQLRepository) Save(name string, price float32) error {
	product := entities.Product{Name: name, Price: price}
	result:= core.BD.Create(&product)
	return result.Error
}

func (repo *MySQLRepository) GetAll() ([]entities.Product, error) {
	var products []entities.Product
	result := core.BD.Find(&products)
	return products, result.Error
}

func (repo *MySQLRepository) Delete(id int32) error {
	result := core.BD.Where ("id_product = ?", id).Delete(&entities.Product{})
    return result.Error
}