package infrastructurem

import (
	"api/src/musics/application"
	"api/src/musics/domain"
	"github.com/gin-gonic/gin"
)

func SetupMusicRoutes(router *gin.Engine, repo domain.IMusic) {
	// Casos de uso
	agregateMusicCaseUse := application.NewAgregateMusic(repo)
	deleteMusicUseCase := application.NewDeleteMusic(repo)
	updateMusicUseCase := application.NewUpdateMusic(repo)
	viewMusicUseCase := application.NewViewMusic(repo)

	// Controladores
	agregateMusicController := NewAgregateMusicController(agregateMusicCaseUse)
	deleteMusicController := NewDeleteMusicController(deleteMusicUseCase)
	updateMusicController := NewUpdateMusicController(updateMusicUseCase)
	viewMusicController := NewViewMusicController(viewMusicUseCase)

	// Rutas
	musicGroup := router.Group("/music")
	{
		musicGroup.POST("", agregateMusicController.Run)
		musicGroup.GET("", viewMusicController.Execute)
		musicGroup.DELETE("/:id", deleteMusicController.Execute)
		musicGroup.PUT("/:id", updateMusicController.Execute)
	}
}
