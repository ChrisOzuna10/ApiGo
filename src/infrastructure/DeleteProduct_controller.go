package infrastructure

import (
	"api/src/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase *application.DeleteProduct
}

func NewDeleteProductController(useCase *application.DeleteProduct) *DeleteProductController {
    return &DeleteProductController{useCase: useCase}
}

func (quit DeleteProductController) Execute(c *gin.Context) {
	var input struct{
		ID int32 `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error": "Invalid product ID"})
	}
	err := quit.useCase.Execute(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"error" : err.Error()})
		return 
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Producto eliminado correctamente"})
}

