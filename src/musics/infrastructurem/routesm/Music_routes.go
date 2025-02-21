package routesm

import (
	"api/src/musics/dependenciesm"
	"github.com/gin-gonic/gin"
)

func SetupMusicRoutes(router *gin.Engine, deps *dependenciesm.DependenciesMusic) {
	musicGroup := router.Group("/musics")

	{
		musicGroup.POST("", deps.CreateMusicController.Run)
		musicGroup.PUT("/:id", deps.UpdateMusicController.Execute)
		musicGroup.GET("", deps.GetAllMusicController)
		musicGroup.GET("/:id", deps.GetByIdController)
		musicGroup.DELETE("/:id", deps.DeleteMusicController.Execute)
	}
}
