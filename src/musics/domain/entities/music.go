package entities

type Music struct {
	ID     int32  `json:"id" gorm:"column:id_music;primaryKey;autoIncrement"`
	Title  string `json:"title"`
	Gender string `json:"gender"`
}
