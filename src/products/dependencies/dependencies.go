package dependencies

import (
	"api/src/products/application/useCaseProduct"
	"api/src/products/domain"
	"api/src/products/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

type ProductDependencies struct {
	CreateProductController  *controllers.CreateProductController
	UpdateProductController  *controllers.UpdateProductController
	DeleteProductController  *controllers.DeleteProductController
	GetAllProductsController gin.HandlerFunc
	GetByIDController        gin.HandlerFunc
}

func NewProductDependencies(repo domain.IProduct) *ProductDependencies {
	createProductUseCase := useCaseProduct.NewCreateProduct(repo)
	updateProductUseCase := useCaseProduct.NewUpdateProduct(repo)
	deleteProductUseCase := useCaseProduct.NewDeleteProduct(repo)

	return &ProductDependencies{
		CreateProductController:  controllers.NewCreateProductController(createProductUseCase),
		UpdateProductController:  controllers.NewUpdateProductController(updateProductUseCase),
		DeleteProductController:  controllers.NewDeleteProductController(deleteProductUseCase),
		GetAllProductsController: controllers.GetAllProductController(repo),
		GetByIDController:        controllers.NewViewByIdProductController(useCaseProduct.NewViewByIdProduct(repo)).Handle,
	}
}
