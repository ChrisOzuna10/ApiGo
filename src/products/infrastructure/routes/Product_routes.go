package routes

import (
	"api/src/products/dependencies"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, deps *dependencies.ProductDependencies) {
	productGroup := router.Group("/product")
	{
		productGroup.POST("", deps.CreateProductController.Run)
		productGroup.PUT("/:id", deps.UpdateProductController.Execute)
		productGroup.GET("", deps.GetAllProductsController)
		productGroup.GET("/:id", deps.GetByIDController)
		productGroup.DELETE("/:id", deps.DeleteProductController.Execute)
	}
}
