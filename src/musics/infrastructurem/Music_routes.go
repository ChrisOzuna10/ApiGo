package infrastructurem

import (
	"api/src/musics/application"
	"api/src/musics/domain"

	"github.com/gin-gonic/gin"
)

func SetupMusicRoutes(router *gin.Engine, repo domain.IMusic) {
	agregateMusicCaseUse := application.NewAgregateMusic(repo)
	agregateMusicController := NewAgregateMusicController(agregateMusicCaseUse)
	deleteMusicUseCase := application.NewDeleteMusic(repo)
	deleteMusicController := NewDeleteMusicController(deleteMusicUseCase)
	updateMusicUseCase := application.NewUpdateMusic(repo)
	updateMusicController := NewUpdateMusicController(updateMusicUseCase)
	musicGroup := router.Group("/music")
	{
		musicGroup.POST("", agregateMusicController.Run)
		musicGroup.GET("", GetAllMusicController(repo))
		musicGroup.DELETE("/:id", deleteMusicController.Execute)
		musicGroup.PUT("/:id", updateMusicController.Execute)
	}
}
