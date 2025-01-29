package infrastructure

import (
	"api/src/products/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllPrductsController(repo domain.IProduct) gin.HandlerFunc {
	return func (c *gin.Context)  {
		products, err := repo.GetAll()
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
		c.JSON (http.StatusOK, products)
	}
}