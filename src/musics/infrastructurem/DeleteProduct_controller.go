package infrastructurem

import (
	"api/src/musics/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteMusicController struct {
	useCase *application.DeleteMusic
}

func NewDeleteMusicController(useCase *application.DeleteMusic) *DeleteMusicController {
	return &DeleteMusicController{useCase: useCase}
}

func (quit DeleteMusicController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Ejecutar el caso de uso para eliminar la música
	err = quit.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Música eliminada correctamente"})
}
