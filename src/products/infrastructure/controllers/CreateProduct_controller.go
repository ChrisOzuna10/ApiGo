package controllers

import (
	"api/src/products/application/useCaseProduct"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductController struct {
	useCase *useCaseProduct.CreateProduct
}

func NewCreateProductController(useCase *useCaseProduct.CreateProduct) *CreateProductController {
	return &CreateProductController{useCase: useCase}
}

func (cp *CreateProductController) Run(c *gin.Context) {
	var input struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cp.useCase.Execute(input.Name, input.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}
