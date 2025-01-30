package infrastructurem

import (
	"api/src/musics/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateMusicController struct {
	useCase *application.UpdateMusic
}

func NewUpdateMusicController(useCase *application.UpdateMusic) *UpdateMusicController {
	return &UpdateMusicController{useCase: useCase}
}

func (cp *UpdateMusicController) Execute(c *gin.Context) {
	// Obtener el ID desde la URL y convertirlo a int32
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Title  string `json:"title"`
		Gender string `json:"gender"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = cp.useCase.Update(int32(id), input.Title, input.Gender)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}
