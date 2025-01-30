package infrastructure

import (
	"api/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase *application.DeleteProduct
}

func NewDeleteProductController(useCase *application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{useCase: useCase}
}

func (quit *DeleteProductController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = quit.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
