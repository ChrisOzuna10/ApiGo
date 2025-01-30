package application

import (
	"api/src/musics/domain"
)

type UpdateMusic struct {
	db domain.IMusic
}

func NewUpdateMusic(db domain.IMusic) *UpdateMusic {
	return &UpdateMusic{db: db}
}

func (up *UpdateMusic) Update(id int32, title string, gender string) error {
	return up.db.Update(id, title, gender)
}
