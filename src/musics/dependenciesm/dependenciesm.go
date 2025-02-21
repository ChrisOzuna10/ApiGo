package dependenciesm

import (
	"api/src/musics/application/useCaseMusic"
	"api/src/musics/domain"
	"api/src/musics/infrastructurem/controllersm"
	"github.com/gin-gonic/gin"
)

type DependenciesMusic struct {
	CreateMusicController *controllersm.AgregateMusicController
	UpdateMusicController *controllersm.UpdateMusicController
	DeleteMusicController *controllersm.DeleteMusicController
	GetAllMusicController gin.HandlerFunc
	GetByIdController     gin.HandlerFunc
}

func NewDependenciesMusic(repo domain.IMusic) *DependenciesMusic {
	createMusicUseCase := useCaseMusic.NewAgregateMusic(repo)
	updateMusicUseCase := useCaseMusic.NewUpdateMusic(repo)
	deleteMusicUseCase := useCaseMusic.NewDeleteMusic(repo)

	return &DependenciesMusic{
		CreateMusicController: controllersm.NewAgregateMusicController(createMusicUseCase),
		UpdateMusicController: controllersm.NewUpdateMusicController(updateMusicUseCase),
		DeleteMusicController: controllersm.NewDeleteMusicController(deleteMusicUseCase),
		GetAllMusicController: controllersm.GetAllMusicController(repo),
		GetByIdController:     controllersm.NewViewByIdMusicController(useCaseMusic.NewViewIdMusic(repo)).Handle,
	}
}
