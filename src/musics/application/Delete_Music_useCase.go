package application

import (
    "api/src/musics/domain"
)

type DeleteMusic struct {
	db  domain.IMusic
}

func NewDeleteMusic(db domain.IMusic) *DeleteMusic {
	return &DeleteMusic{db: db}
}

func (db *DeleteMusic) Execute(id int32) error {
	return db.db.Delete(id)
}