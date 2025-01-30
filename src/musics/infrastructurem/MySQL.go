package infrastructurem

import (
	"api/src/core"
	"api/src/musics/domain"
	"api/src/musics/domain/entities"
)

type MySQLRepositoryMusics struct{}

func NewMySQLRepositoryMusics() domain.IMusic {
	return &MySQLRepositoryMusics{}
}

func (repo *MySQLRepositoryMusics) Save(title string, gender string) error {
	music := entities.Music{Title: title, Gender: gender}
	result := core.BD.Create(&music)
	return result.Error
}
func (repo *MySQLRepositoryMusics) GetAll() ([]entities.Music, error) {
	var musics []entities.Music
	result := core.BD.Find(&musics)
	return musics, result.Error
}

func (repo *MySQLRepositoryMusics) Delete(id int32) error {
	result := core.BD.Where("id_music = ?", id).Delete(&entities.Music{})
	return result.Error
}

func (repo *MySQLRepositoryMusics) Update(id int32, title string, gender string) error {
	result := core.BD.Model(&entities.Music{}).Where("id_music = ?", id).
		Updates(entities.Music{Title: title, Gender: gender})
	return result.Error
}
