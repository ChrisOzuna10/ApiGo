package infrastructure

import (
	"api/src/core"
	"api/src/products/domain"
	"api/src/products/domain/entities"
	"fmt"
)

type MySQLRepositoryProducts struct{}

func NewMySQLRepositoryProducts() domain.IProduct {
	return &MySQLRepositoryProducts{}
}

func (repo *MySQLRepositoryProducts) Save(name string, price float32) error {
	product := entities.Product{Name: name, Price: price}
	result := core.BD.Create(&product)
	return result.Error
}

func (repo *MySQLRepositoryProducts) GetAll() ([]entities.Product, error) {
	var products []entities.Product
	result := core.BD.Find(&products)
	return products, result.Error
}

func (repo *MySQLRepositoryProducts) Delete(id int32) error {
	result := core.BD.Where("id_product = ?", id).Delete(&entities.Product{})
	return result.Error
}

func (repo *MySQLRepositoryProducts) Update(id int32, name string, price float32) error {
	result := core.BD.Model(&entities.Product{}).Where("id_product = ?", id).
		Updates(entities.Product{Name: name, Price: price})
	return result.Error
}

func (repo *MySQLRepositoryProducts) GetById(id int32) (*entities.Product, error) {
	var product entities.Product
	result := core.BD.First(&product, "id_product = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		fmt.Println("No se encontró ningún producto con id:", id)
		return nil, fmt.Errorf("producto no encontrado")
	}
	fmt.Println("Producto encontrado", product)
	return &product, nil
}
