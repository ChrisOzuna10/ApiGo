package controllersm

import (
	"api/src/musics/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllMusicController(repo domain.IMusic) gin.HandlerFunc {
	return func(c *gin.Context) {
		music, err := repo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, music)
	}
}
