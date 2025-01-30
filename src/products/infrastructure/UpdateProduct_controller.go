package infrastructure

import (
	"api/src/products/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateProductController struct {
	useCase *application.UpdateProduct
}

func NewUpdateProductController(useCase *application.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{useCase: useCase}
}

func (cp *UpdateProductController) Execute(c *gin.Context) {
	// Obtener el ID desde la URL y convertirlo a int32
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = cp.useCase.Update(int32(id), input.Name, input.Price)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}
