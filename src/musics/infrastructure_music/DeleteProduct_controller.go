package infrastructuremusic

import (
    "api/src/musics/application"
    "net/http"

	"github.com/gin-gonic/gin"
)

type DeleteMusicController struct {
	useCase *application.DeleteMusic
}

func NewDeleteMusicController(useCase *application.DeleteMusic) *DeleteMusicController {
    return &DeleteMusicController{useCase: useCase}
}

func (quit DeleteMusicController) Execute(c *gin.Context) {
	var input struct {
		ID int32 `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	err := quit.useCase.Execute(input.ID)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusCreated, gin.H{"message": "Música eliminada correctamente"})
}