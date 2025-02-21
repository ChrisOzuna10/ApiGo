package useCaseMusic

import (
	"api/src/musics/domain"
	"api/src/musics/domain/entities"
)

type ViewByIdMusic struct {
	db domain.IMusic
}

func NewViewIdMusic(db domain.IMusic) *ViewByIdMusic {
	return &ViewByIdMusic{db: db}
}

func (vp *ViewByIdMusic) GetById(id int32) (*entities.Music, error) {
	return vp.db.GetById(id)
}
