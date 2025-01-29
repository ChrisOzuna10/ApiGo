package application

import ( "api/src/musics/domain"
)

type AgregateMusic struct {
	db domain.IMusic
}

func NewAgregateMusic( db domain.IMusic) *AgregateMusic {
	return &AgregateMusic{db: db}	
}

func (cp *AgregateMusic) Execute(tittle string, gender string) error {
    return cp.db.Save(tittle, gender)
}
