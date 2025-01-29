package domain

import "api/src/musics/domain/entities"

type IMusic interface {
    Save(tittle string, gender string) error
    GetAll() ([]entities.Music, error)
    Delete(id int32) error
}