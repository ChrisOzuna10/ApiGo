package infrastructure

import (
	"api/src/products/application"
	"api/src/products/domain"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes (router *gin.Engine, repo domain.IProduct){
	createProductCaseUse := application.NewCreateProduct(repo)
	createProductController := NewCreateProductController(createProductCaseUse)
	deleteProductUseCase := application.NewDeleteProduct(repo)
	deleteProductController := NewDeleteProductController(deleteProductUseCase)
	productGroup := router.Group("/products")
	{
        productGroup.POST("", createProductController.Run)
		productGroup.GET("", GetAllPrductsController(repo))
		productGroup.DELETE("", deleteProductController.Execute)
	}
}