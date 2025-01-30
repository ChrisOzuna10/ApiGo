package application

import (
	"api/src/musics/domain"
	"api/src/musics/domain/entities"
)

type ViewMusic struct {
	db domain.IMusic
}

func NewViewMusic(db domain.IMusic) *ViewMusic {
	return &ViewMusic{db: db}
}

func (vp *ViewMusic) Execute() ([]entities.Music, error) {
	return vp.db.GetAll()
}
