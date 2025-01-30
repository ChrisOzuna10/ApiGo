package infrastructure

import (
	"api/src/products/application"
	"api/src/products/domain"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, repo domain.IProduct) {
	createProductCaseUse := application.NewCreateProduct(repo)
	createProductController := NewCreateProductController(createProductCaseUse)
	deleteProductUseCase := application.NewDeleteProduct(repo)
	deleteProductController := NewDeleteProductController(deleteProductUseCase)
	updateProductUseCase := application.NewUpdateProduct(repo)
	updateProductController := NewUpdateProductController(updateProductUseCase)
	productGroup := router.Group("/products")
	{
		productGroup.POST("", createProductController.Run)
		productGroup.GET("", GetAllPrductsController(repo))
		productGroup.DELETE("/:id", deleteProductController.Execute)
		productGroup.PUT("/:id", updateProductController.Execute)
	}
}
