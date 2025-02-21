package controllers

import (
	"api/src/products/application/useCaseProduct"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

type ViewByIdProductController struct {
	useCase *useCaseProduct.ViewByIdProduct
}

func NewViewByIdProductController(useCase *useCaseProduct.ViewByIdProduct) *ViewByIdProductController {
	return &ViewByIdProductController{useCase: useCase}
}

func (controller *ViewByIdProductController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"message": "ID invalido"})
		return
	}
	product, err := controller.useCase.GetById(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto encontrado", "Product": product})
}
