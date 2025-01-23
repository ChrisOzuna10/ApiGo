package infrastructure

import (
	"api/src/application"
	"api/src/domain"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes (router *gin.Engine, repo domain.IProduct){
	createProductCaseUse := application.NewCreateProduct(repo)
	createProductController := NewCreateProductController(createProductCaseUse)

	productGroup := router.Group("/products")
	{
        productGroup.POST("", createProductController.Run)
		productGroup.GET("", GetAllPrductsController(repo))
	}
}