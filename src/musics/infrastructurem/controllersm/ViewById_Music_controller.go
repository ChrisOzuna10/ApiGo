package controllersm

import (
	"api/src/musics/application/useCaseMusic"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

type ViewByIdMusicController struct {
	useCase *useCaseMusic.ViewByIdMusic
}

func NewViewByIdMusicController(useCse *useCaseMusic.ViewByIdMusic) *ViewByIdMusicController {
	return &ViewByIdMusicController{useCase: useCse}
}

func (controller *ViewByIdMusicController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	music, err := controller.useCase.GetById(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"music": music})
}
