package domain

import "api/src/musics/domain/entities"

type IMusic interface {
	Save(title string, gender string) error
	GetAll() ([]entities.Music, error)
	GetById(id int32) (*entities.Music, error)
	Delete(id int32) error
	Update(id int32, title string, genre string) error
}
