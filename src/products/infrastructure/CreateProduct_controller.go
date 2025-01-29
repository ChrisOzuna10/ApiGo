package infrastructure

import (
	"api/src/products/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	useCase *application.CreateProduct
}

func NewCreateProductController(useCase *application.CreateProduct) *CreateProductController {
	return &CreateProductController{useCase: useCase}
}

func (cp * CreateProductController) Run(c * gin.Context) {
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