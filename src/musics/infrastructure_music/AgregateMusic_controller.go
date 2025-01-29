package infrastructuremusic

import (
	"api/src/musics/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AgregateMusicController struct {
	useCase *application.AgregateMusic
}

func NewAgregateMusicController(useCase *application.AgregateMusic) *AgregateMusicController {
    return &AgregateMusicController{useCase: useCase}
}

func (cp * AgregateMusicController) Run(c * gin.Context) {
	var input struct{
		Title string `json:"title"`
        Gender string `json:"gender"`
	}

if err := c.ShouldBindJSON(&input); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

err := cp.useCase.Execute(input.Title, input.Gender)
if err!= nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
}

c.JSON(http.StatusCreated, gin.H{"message": "Agregado correctamente"})
}