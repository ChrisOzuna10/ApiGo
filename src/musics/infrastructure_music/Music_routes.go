package infrastructuremusic

import (
	"api/src/musics/application"
	"api/src/musics/domain"

	"github.com/gin-gonic/gin"
)

func SetupMusicRoutes (router *gin.Engine, repo domain.IMusic) {
	agregateMusicCaseUse := application.NewAgregateMusic(repo)
	agregateMusicController := NewAgregateMusicController(agregateMusicCaseUse)
	deleteMusicUseCase := application.NewDeleteMusic(repo)
	deleteMusicController := NewDeleteMusicController(deleteMusicUseCase)
	musicGroup := router.Group("/music")
	{
        musicGroup.POST("", agregateMusicController.Run)
		musicGroup.GET("", GetAllMusicController(repo))
		musicGroup.DELETE("", deleteMusicController.Execute)
    }
}